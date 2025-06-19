package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Order struct {
	ID        string    `json:"id"`
	UserID    int64     `json:"user_id"`
	PetID     int64     `json:"pet_id"`
	Status    string    `json:"status"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
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

	r.POST("/orders", createOrder)
	r.PUT("/orders/:id/pay", payOrder)

	fmt.Println("订单服务启动中，监听端口:8084...")
	r.Run(":8084")
}

func createOrder(c *gin.Context) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 记录传入的user_id
	log.Printf("Received user_id: %d", order.UserID)

	// 验证用户是否存在
	var count int64
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", order.UserID).Scan(&count)
	if err != nil {
		log.Printf("Error verifying user existence for user_id %d: %v", order.UserID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify user existence"})
		return
	}
	if count == 0 {
		log.Printf("User with ID %d does not exist", order.UserID)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("User with ID %d does not exist. Please register first.", order.UserID)})
		return
	}

	// 设置默认值
	if order.StartDate.IsZero() {
		order.StartDate = time.Now()
	}
	if order.EndDate.IsZero() {
		order.EndDate = time.Now().Add(24 * time.Hour)
	}

	tx, err := db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		return
	}

	tableName := fmt.Sprintf("order_%d", hashOrderID(order.UserID))
	insertSQL := "INSERT INTO " + tableName + " (user_id, pet_id, status, start_date, end_date) VALUES(?, ?, ?, ?, ?)"
	log.Printf("Executing SQL: %s with params: %d, %d, %s, %s, %s", insertSQL, order.UserID, order.PetID, "created", order.StartDate, order.EndDate)
	result, err := tx.Exec(
		insertSQL,
		order.UserID, order.PetID, "created", order.StartDate, order.EndDate,
	)
	if err != nil {
		log.Printf("Error inserting into table %s: %v", tableName, err)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get order ID"})
		return
	}

	order.ID = strconv.FormatInt(orderID, 10)
	order.Status = "created"

	err = tx.Commit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to complete order creation"})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func hashOrderID(id int64) int64 {
	return id % 8
}

func payOrder(c *gin.Context) {
	id := c.Param("id")
	orderID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	tableNum := hashOrderID(orderID)
	tableName := fmt.Sprintf("order_%d", tableNum)

	result, err := db.Exec("UPDATE "+tableName+" SET status = ? WHERE id = ?",
		"paid", orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify update"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	var order Order
	err = db.QueryRow("SELECT id, user_id, pet_id, status FROM "+tableName+" WHERE id = ?",
		orderID).Scan(&order.ID, &order.UserID, &order.PetID, &order.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated order"})
		return
	}

	c.JSON(http.StatusOK, order)
}
