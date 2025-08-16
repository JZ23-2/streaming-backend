package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func StreamRoutes(api *gin.RouterGroup) {
	stream := api.Group("/streams")
	{
		stream.POST("/start")
		stream.POST("/create-stream", controllers.CreateStreamController)
	}
}
