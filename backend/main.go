package main

import (
	"bytes"
	"fmt"
	"gin-storage/routes"
	"io"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(CORSMiddleware())
	// router.Use(RequestLoggerMiddleware())
	
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	routes.ProfileRoutes(router)
	routes.FileRoutes(router)
	// r.POST("/v1/form", controllers.HandleFileUploadToBucket)
	router.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("request is accepted")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length,Accept,Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Content-Language, Accept-Language")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,DELETE")
	
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
	
		c.Next()
    }
}


func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

		fmt.Println(readBody(rdr1)) // Print request body
		c.Request.Body = rdr2
		c.Next()
	}
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}