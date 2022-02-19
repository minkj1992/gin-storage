package routes

import (
	"gin-storage/controllers"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(incomingRoutes *gin.Engine) {
	profiles := incomingRoutes.Group("v1/profiles")
	{
		profiles.GET("/", controllers.Ping)
		profiles.POST("/", controllers.CreateProfile)
	}
}