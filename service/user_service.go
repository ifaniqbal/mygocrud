package service

import (
	"fmt"
	"mygocrud/model"
	"mygocrud/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadUsersHandler(c *gin.Context) {
	var users []model.User

	query := `select * from users`
	filter := c.Query("filter")
	if filter != "" {
		query = fmt.Sprintf(
			"%s %s",
			query,
			"where username = ?",
		)
	}

	err := repository.Db.Raw(query, filter).Scan(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Message: fmt.Sprintf("failed to save user: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Success",
		Data:    users,
	})
}

func CreateUserHandler(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())),
		)
		return
	}

	err = repository.Db.Create(&user).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to save user: %s", err.Error())),
		)
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse("Success", user))
}

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
	var product model.Product
	err := c.ShouldBind(&product)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())),
		)
		return
	}

	err = repository.Db.Create(&product).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to save product: %s", err.Error())),
		)
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse("Success", product))
}
