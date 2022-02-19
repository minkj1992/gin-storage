package routes

import (
	"gin-storage/controllers"

	"github.com/gin-gonic/gin"
)

func FileRoutes(incomingRoutes *gin.Engine) {
	files := incomingRoutes.Group("v1/files")
	{
		files.GET("/", controllers.HandleFileDownloadFromBucket)
		files.POST("/", controllers.HandleFileUploadToBucket)
	}
}