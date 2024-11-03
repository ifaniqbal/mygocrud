package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseRequest struct {
	Page     int    `json:"page,omitempty" form:"page"`
	PageSize int    `json:"page_size,omitempty" form:"page_size"`
	Filter   string `json:"filter,omitempty" form:"filter"`
}

type BaseResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

var items []Item

type Item struct {
	Name        string `json:"name"`
	Code        int    `json:"code"`
	PostingDate string `json:"posting_date"`
}

func main() {
	r := gin.Default()

	// Routes
	r.POST("/items", CreateItemHandler)
	r.GET("/items", ReadItemsHandler)

	// Start server
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func CreateItemHandler(c *gin.Context) {
	var newItem Item
	err := c.ShouldBind(&newItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, BaseResponse{
			Success: false,
			Data:    nil,
			Message: fmt.Sprintf("failed to bind request: %s", err.Error()),
		})
		return
	}

	items = append(items, newItem)

	c.JSON(http.StatusOK, BaseResponse{
		Success: true,
		Message: "Success",
		Data:    newItem,
	})
}

func ReadItemsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, BaseResponse{
		Success: true,
		Message: "Success",
		Data:    items,
	})
}
