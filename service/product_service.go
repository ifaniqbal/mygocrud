package service

import (
	"fmt"
	"mygocrud/model"
	"mygocrud/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ReadProductsHandler(c *gin.Context) {
	var products []model.Product

	query := `select * from products`
	filter := c.Query("filter")
	var args []any
	if filter != "" {
		query = fmt.Sprintf(
			"%s %s",
			query,
			"where name = ?",
		)

		args = append(args, filter)
	}

	err := repository.Db.Raw(query, args...).Scan(&products).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Message: fmt.Sprintf("failed to save product: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Success",
		Data:    products,
	})
}

func CreateProductHandler(c *gin.Context) {
	var productDto model.ProductDto
	err := c.ShouldBind(&productDto)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())),
		)
		return
	}

	product := model.Product{
		ID:        productDto.ID,
		Name:      productDto.Name,
		Price:     productDto.Price,
		Stock:     productDto.Stock,
		CreatedAt: productDto.CreatedAt,
		UpdatedAt: productDto.UpdatedAt,
	}
	if productDto.Description != nil {
		product.Description.String = *productDto.Description
		product.Description.Valid = true
	}

	err = repository.Db.Create(&product).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to save product: %s", err.Error())),
		)
		return
	}

	productDto.ID = product.ID
	productDto.CreatedAt = product.CreatedAt
	productDto.UpdatedAt = product.UpdatedAt

	c.JSON(http.StatusOK, model.NewSuccessResponse("Success", productDto))
}

func UpdateProductHandler(c *gin.Context) {
	var productDto model.ProductDto
	err := c.ShouldBind(&productDto)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())),
		)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("invalid id: %s", err.Error())),
		)
		return
	}

	product := model.Product{
		ID:        id,
		Name:      productDto.Name,
		Price:     productDto.Price,
		Stock:     productDto.Stock,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	if productDto.Description != nil {
		product.Description.String = *productDto.Description
		product.Description.Valid = true
	}

	err = repository.Db.Save(&product).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to update product: %s", err.Error())),
		)
		return
	}

	productDto.ID = product.ID
	productDto.CreatedAt = product.CreatedAt
	productDto.UpdatedAt = product.UpdatedAt

	c.JSON(http.StatusOK, model.NewSuccessResponse("Success", productDto))
}
