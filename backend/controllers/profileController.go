package controllers

import (
	"fmt"
	"log"
	"mime/multipart"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ProfileForm struct {
	Files []*multipart.FileHeader `form:"files" binding:"required"`
	Data Profile                  `form:"data" binding:"required"`
}

type Profile struct {
	FirstName   string `json:"firstName" validate:"required"`
	LastName   string `json:"lastName" validate:"required"`
	Email 	*string 	`json:"email" validate:"email,required"`
	HasPhone bool 	`json:hasPhone`
	PhoneNumber string `json:phoneNumber`
}


func Ping(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
}


func CreateProfile(c *gin.Context) {
	obj := ProfileForm{}
	if !strings.Contains(c.GetHeader("Content-Type"), "multipart") {
		fmt.Println(c.GetHeader("Content-Type"))
		c.JSON(401, gin.H{"test": "test"})
		return
	}
	// Multipart form
	if err := c.BindWith(&obj, binding.FormMultipart); err != nil {
		c.JSON(406, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("%+v\n", obj)

	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	fmt.Printf("%+v\n", form)
	fmt.Printf("%+v\n", files)

	for _, file := range files {
		log.Println(file.Filename)

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, "./files")
	}


	
	// c.BindJSON(&profile)
	// services.HandleFileUploadToBucket(c)
}

