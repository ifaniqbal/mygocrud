package main

import (
	"fmt"
	"log"
	"mygocrud/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var items []model.Item

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
	r.POST("/items", CreateItemHandler)
	r.GET("/items", ReadItemsHandler)

	// Start server
	err = r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func CreateItemHandler(c *gin.Context) {
	var newItem model.Item
	err := c.ShouldBind(&newItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BaseResponse{
			Success: false,
			Data:    nil,
			Message: fmt.Sprintf("failed to bind request: %s", err.Error()),
		})
		return
	}

	items = append(items, newItem)

	c.JSON(http.StatusOK, model.BaseResponse{
		Success: true,
		Message: "Success",
		Data:    newItem,
	})
}

func ReadItemsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.BaseResponse{
		Success: true,
		Message: "Success",
		Data:    items,
	})
}
