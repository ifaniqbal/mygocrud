package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mygocrud/model"
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

	authGroup := r.Group("/")
	// gin auth
	// authorization
	// next
	// after middleware
	// log gin auth
	authGroup.Use(ginAuth())
	authGroup.Use(ginLoggingAfterMiddleware())
	authGroup.POST("/products", service.CreateProductHandler)
	authGroup.GET("/products", service.ReadProductsHandler)
	authGroup.GET("/products/:id", service.ReadByIdProductsHandler)
	authGroup.PUT("/products/:id", service.UpdateProductHandler)
	authGroup.DELETE("/products/:id", service.DeleteProductHandler)
	authGroup.POST("/upload-product-image", service.UploadProductImageHandler)

	//Start server
	err = r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
		return
	}

	//http.Handle("/upload-product-image", afterMiddleware(
	//	loggingBeforeMiddleware(authorizationMiddleware(http.HandlerFunc(uploadProductImageV2))),
	//	loggingAfterMiddleware,
	//))

	//http.HandleFunc("/upload-product-image", uploadProductImageV2)
	//
	//err = http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	log.Fatalln(err)
	//	return
	//}
}

func ginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		if auth == "" {
			c.String(http.StatusUnauthorized, "empty token")
			c.Abort()
			return
		}

		c.Next()

		log.Println("Gin Auth:", c.Request.URL)
	}
}

func ginLoggingAfterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		log.Println("After Middleware:", c.Request.URL)
	}
}

func authorize(w http.ResponseWriter, r *http.Request) bool {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		http.Error(w, "empty token", http.StatusUnauthorized)
		return false
	}

	return true
}

func authorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "empty token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// afterMiddleware wraps a handler and performs tasks after it runs.
func afterMiddleware(handler http.Handler, afterFunc func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Run the main handler
		handler.ServeHTTP(w, r)

		// Run the after middleware function
		afterFunc(w, r)
	})
}

func loggingBeforeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func loggingAfterMiddleware(w http.ResponseWriter, r *http.Request) {
	log.Println("Response:", r.Method, r.URL.Path)
}

func uploadProductImageV2(w http.ResponseWriter, r *http.Request) {
	ok := authorize(w, r)
	if !ok {
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	r.URL.Query()

	//auth := r.Header.Get("Authorization")
	//fmt.Println(auth)

	formFile, file, err := r.FormFile("image")
	if err != nil {
		_, _ = fmt.Fprintln(w)
		http.Error(
			w,
			fmt.Sprintf("failed to get uploaded file: %s", err.Error()),
			http.StatusBadRequest,
		)
		return
	}
	defer formFile.Close()

	maxFileSize := 100 << 10 // 10kb
	if file.Size > int64(maxFileSize) {
		http.Error(
			w,
			"file size exceed maximum allowed",
			http.StatusBadRequest,
		)
		return
	}

	var productDto model.ProductDto
	productDto.Name = "tes"
	res, err := json.Marshal(productDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(res)
}
