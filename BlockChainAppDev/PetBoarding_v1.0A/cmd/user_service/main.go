package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

var db *sql.DB

func initDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/pet_boarding?parseTime=true&loc=Local")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	var err error
	db, err = initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()

	r.POST("/register", register)
	r.POST("/login", login)

	fmt.Println("用户服务启动中...")
	r.Run(":8081")
}

func register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", user.Username).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "检查用户名失败"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名已存在"})
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Printf("Failed to begin transaction: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法启动交易"})
		return
	}

	_, err = tx.Exec("INSERT INTO users(username, password, phone, email) VALUES(?, ?, ?, ?)", user.Username, user.Password, user.Phone, user.Email)
	if err != nil {
		tx.Rollback()
		log.Printf("插入用户失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册用户失败"})
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("提交事务失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册用户失败"})
		return
	}

	log.Printf("用户注册成功: %s", user.Username)
	c.JSON(http.StatusCreated, gin.H{"message": "用户注册成功"})
}

func login(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", user.Username).Scan(&storedPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效凭据"})
		return
	}
	if storedPassword != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效凭据"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
}
