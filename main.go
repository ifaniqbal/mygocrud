package main

import (
	"log"
	"mygocrud/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// Database connection
	dsn := "root:@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	r := gin.Default()

	// Routes
	r.POST("/items", service.CreateItemHandler)
	r.GET("/items", service.ReadItemsHandler)

	// Start server
	err = r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
