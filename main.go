package main

import (
	"log"
	"mygocrud/repository"
	"mygocrud/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Database connection
	dsn := "root:@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	repository.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	r := gin.Default()

	// Routes
	r.POST("/items", service.CreateItemHandler)
	r.GET("/items", service.ReadItemsHandler)

	r.POST("/users", service.CreateUserHandler)
	r.GET("/users", service.ReadUsersHandler)

	r.POST("/products", service.CreateProductHandler)
	r.GET("/products", service.ReadProductsHandler)
	r.GET("/products/:id", service.ReadByIdProductsHandler)
	r.PUT("/products/:id", service.UpdateProductHandler)
	r.DELETE("/products/:id", service.DeleteProductHandler)
	r.POST("/upload-product-image", service.UploadProductImageHandler)

	// Start server
	err = r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
