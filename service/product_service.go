package service

import (
	"errors"
	"fmt"
	"mygocrud/model"
	"mygocrud/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
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

func ReadByIdProductsHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("invalid id: %s", idParam)),
		)
		return
	}

	var product model.Product
	err = repository.Db.First(&product, id).Error
	statusCodeError := http.StatusInternalServerError
	if errors.Is(err, gorm.ErrRecordNotFound) {
		statusCodeError = http.StatusNotFound
	}

	if err != nil {
		c.JSON(statusCodeError, model.Response{
			Message: fmt.Sprintf("failed to get product: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Success",
		Data:    product,
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
		Price:     decimal.NewFromInt(int64(productDto.Price)),
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

	// validasi logika

	var product model.Product
	err = repository.Db.First(&product, id).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("invalid id: %s", err.Error())),
		)
		return
	}

	product.Name = productDto.Name
	product.Price = decimal.NewFromInt(int64(productDto.Price))
	product.Stock = productDto.Stock
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

func DeleteProductHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("invalid id: %s", err.Error())),
		)
		return
	}

	product := model.Product{ID: id}
	result := repository.Db.Delete(product)
	if result.Error != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to delete product: %s", result.Error.Error())),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		model.NewSuccessResponse(fmt.Sprintf("%d products deleted", result.RowsAffected), nil),
	)
}
