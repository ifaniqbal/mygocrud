package main

import (
	"fmt"
	"log"
	"mygocrud/repository"
	"mygocrud/service"
	"net/http"

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
	r.MaxMultipartMemory = 12 << 10 // 12kb

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
	//err = r.Run(":8080")
	//if err != nil {
	//	log.Fatalln(err)
	//	return
	//}

	http.HandleFunc("/upload-product-image", func(w http.ResponseWriter, r *http.Request) {
		formFile, file, err := r.FormFile("image")
		if err != nil {
			_, _ = fmt.Fprintln(w, fmt.Sprintf("failed to get uploaded file: %s", err.Error()))
			return
		}
		defer formFile.Close()

		maxFileSize := 100 << 10 // 10kb
		if file.Size > int64(maxFileSize) {
			_, _ = fmt.Fprintln(w, "file size exceed maximum allowed")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintln(w, `{"message": "Success"}`)
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
