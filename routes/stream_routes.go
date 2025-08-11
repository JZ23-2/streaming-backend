package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func StreamRoutes(api *gin.RouterGroup) {
	stream := api.Group("/stream")
	{
		stream.POST("/publish", controllers.HandlePublish)
		stream.POST("/view", controllers.HandleView)
	}
}
