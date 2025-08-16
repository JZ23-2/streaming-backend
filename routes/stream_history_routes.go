package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func StreamHistoryRoutes(api *gin.RouterGroup) {
	stream := api.Group("/stream-history")
	{
		stream.POST("/create", controllers.CreateStreamHistoryController)
		stream.GET("/all-stream", controllers.GetAllStreamHistoryByStreamerIDController)
		stream.GET("/by-id", controllers.GetAllStreamHistoryByIdController)
	}
}
