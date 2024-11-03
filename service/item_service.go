package service

import (
	"fmt"
	"mygocrud/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

var items []model.Item
