package main

import (
	"fmt"
	"log"
	"mygocrud/model"
	"mygocrud/service"
	"net/http"

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

	r.POST("/users", CreateUserHandler)
	r.GET("/users", ReadUsersHandler)

	// Start server
	err = r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
		return
	}
}

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

	// select * from users where username = ?

	err := db.Raw(query, filter).Scan(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.BaseResponse{
			Message: fmt.Sprintf("failed to save user: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, model.BaseResponse{
		Success: true,
		Message: "Success",
		Data:    users,
	})
}

func CreateUserHandler(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BaseResponse{
			Message: fmt.Sprintf("failed to bind request: %s", err.Error()),
		})
		return
	}

	err = db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.BaseResponse{
			Message: fmt.Sprintf("failed to save user: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, model.BaseResponse{
		Success: true,
		Message: "Success",
		Data:    user,
	})
}
