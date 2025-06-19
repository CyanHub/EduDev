package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Review struct {
	ID        string  `json:"id"`
	Content   string  `json:"content"`
	Rating    float64 `json:"rating"`
	UserID    int64   `json:"user_id"`
	ServiceID int64   `json:"service_id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
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

	r.POST("/reviews", createReview)
	r.GET("/reviews/:id", getReview)
	r.PUT("/reviews/:id", updateReview)
	r.DELETE("/reviews/:id", deleteReview)

	fmt.Println("评价服务启动中，监听端口:8083...")
	r.Run(":8083")
}

func createReview(c *gin.Context) {
	var review Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO providers (user_id, service_type, rating) VALUES (?, ?, ?)",
		review.UserID, review.Content, review.Rating)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	review.ID = fmt.Sprintf("review_%d", id)
	c.JSON(http.StatusCreated, review)
}

func getReview(c *gin.Context) {
	id := c.Param("id")
	var review Review
	err := db.QueryRow("SELECT id, user_id, service_type as content, rating FROM providers WHERE id = ?", id).Scan(
		&review.ID, &review.UserID, &review.Content, &review.Rating)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到评价"})
		return
	}

	c.JSON(http.StatusOK, review)
}

func updateReview(c *gin.Context) {
	id := c.Param("id")
	var review Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("UPDATE providers SET service_type = ?, rating = ? WHERE id = ?",
		review.Content, review.Rating, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, review)
}

func deleteReview(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM providers WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "评价删除成功"})
}
