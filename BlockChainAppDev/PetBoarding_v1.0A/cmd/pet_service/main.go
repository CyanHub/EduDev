package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Pet struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	OwnerID int64 `json:"owner_id"`
	Age     int    `json:"age,omitempty"`
}

var db *sql.DB

func main() {
	// 初始化数据库连接
	initDB()

	r := gin.Default()

	r.POST("/pets", createPet)
	r.GET("/pets/:id", getPet)
	r.PUT("/pets/:id", updatePet)
	r.DELETE("/pets/:id", deletePet)

	fmt.Println("宠物服务启动中，监听端口:8082...")
	r.Run(":8082")
}

func initDB() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/pet_boarding?parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func createPet(c *gin.Context) {
	var pet struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Type    string `json:"type"`
		OwnerID int64 `json:"owner_id"`
		Age     int    `json:"age,omitempty"`
	}
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Printf("无法开始交易: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法启动交易"})
		return
	}

	result, err := tx.Exec("INSERT INTO pet_basic(name, type, owner_id, age) VALUES(?, ?, ?, ?)",
		pet.Name, pet.Type, pet.OwnerID, pet.Age)
	if err != nil {
		tx.Rollback()
		log.Printf("无法插入宠物: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建宠物失败"})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Printf("获取最后插入ID失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建宠物失败"})
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("交易提交失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "未能完成宠物创建"})
		return
	}

	pet.ID = fmt.Sprintf("pet_%d", id)
	log.Printf("成功创建宠物: %s", pet.ID)
	c.JSON(http.StatusCreated, pet)
}

func getPet(c *gin.Context) {
	id := c.Param("id")
	var pet Pet

	err := db.QueryRow("SELECT id, name, type, owner_id, age FROM pet_basic WHERE id = ?", id).Scan(
		&pet.ID, &pet.Name, &pet.Type, &pet.OwnerID, &pet.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到宠物"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, pet)
}

func updatePet(c *gin.Context) {
	id := c.Param("id")
	var pet Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Printf("交易启动失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "未能启动交易"})
		return
	}

	result, err := tx.Exec("UPDATE pet_basic SET name = ?, type = ?, owner_id = ?, age = ? WHERE id = ?",
		pet.Name, pet.Type, pet.OwnerID, pet.Age, id)
	if err != nil {
		tx.Rollback()
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到宠物"})
		} else {
			log.Printf("更新宠物失败: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "更新宠物失败"})
		}
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		log.Printf("获取受影响行数失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新宠物失败"})
		return
	}

	if rowsAffected == 0 {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到宠物"})
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("交易提交失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "未能完成宠物更新"})
		return
	}

	pet.ID = id
	log.Printf("成功更新了宠物: %s", pet.ID)
	c.JSON(http.StatusOK, pet)
}

func deletePet(c *gin.Context) {
	id := c.Param("id")

	tx, err := db.Begin()
	if err != nil {
		log.Printf("交易启动失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "未能启动交易"})
		return
	}

	result, err := tx.Exec("DELETE FROM pet_basic WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		log.Printf("删除宠物失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除宠物失败"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		log.Printf("获取受影响行数失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除宠物失败"})
		return
	}

	if rowsAffected == 0 {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到宠物"})
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("交易提交失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "未能完成宠物删除"})
		return
	}

	log.Printf("成功删除宠物: %s", id)
	c.JSON(http.StatusOK, gin.H{"message": "宠物删除成功"})
}
