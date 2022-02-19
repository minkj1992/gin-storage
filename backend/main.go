package main

import (
	"fmt"
	"gin-storage/controllers"

	"github.com/gin-gonic/gin"
)


func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/files/", controllers.HandleFileUploadToBucket)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")	
}