package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	//dirName := "DIRECTORY"
	//newDirName := "NEW_DIRECTORY"
	//fileName := "FILE"
	//newFileName := "NEW_FILE"
	//filePath := filepath.Join(dirName, fileName)
	//newFilePath := filepath.Join(dirName, newFileName)
	//newFileDirPath := filepath.Join(newDirName, newFileName)

	//createDir(dirName)
	//createFile(filePath, "hello world\n")
	//appendToFile(filePath, "it works\n")
	//readFile(filePath)
	//renameFile(filePath, newFileDirPath)
	//deleteFile(filePath)
	//deleteDir(dirName)

	r := gin.Default()

	uploadDir := "UPLOAD"

	r.POST("/upload", func(c *gin.Context) {
		multipartFileHeader, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve file"})
			return
		}

		name := c.PostForm("name")
		path := filepath.Join(uploadDir, name)
		if err = c.SaveUploadedFile(multipartFileHeader, path); err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": fmt.Sprintf("%s: %s", "Failed to save file", err.Error())},
			)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Success", "filename": multipartFileHeader.Filename})
	})

	//r.GET("/download", func(c *gin.Context) {
	//	name := c.Query("file")
	//	path := filepath.Join("DIRECTORY", name)
	//	if _, err := filepath.Abs(path); err != nil {
	//		c.Status(http.StatusNotFound)
	//		return
	//	}
	//
	//	c.FileAttachment(path, name)
	//})
	//
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func createDir(name string) {
	err := os.Mkdir(name, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	fmt.Println("Directory created successfully")
}

func createFile(name string, content string) {
	// Create a new file or open if it exists
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("failed to close file:", err.Error())
		}
	}()

	// Write to the file
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File created and data written successfully")
}

func appendToFile(name string, content string) {
	// Create a new file or open if it exists
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("failed to close file:", err.Error())
		}
	}()

	// Write to the file
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File created and data written successfully")
}

func readFile(name string) {
	// Read file contents
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File contents:", string(data))
}

func renameFile(from, to string) {
	// Rename or move the file
	err := os.Rename(from, to)
	if err != nil {
		fmt.Println("Error renaming file:", err)
		return
	}

	fmt.Println("File renamed successfully")
}

func deleteFile(name string) {
	err := os.Remove(name)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}

	fmt.Println("File deleted successfully")
}

func deleteDir(name string) {
	// Delete a directory (must be empty to delete)
	err := os.Remove(name)
	if err != nil {
		fmt.Println("Error deleting directory:", err)
		return
	}

	fmt.Println("Directory deleted successfully")
}
