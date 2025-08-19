package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func StreamRoutes(api *gin.RouterGroup) {
	stream := api.Group("/streams")
	{
		stream.POST("/create-stream", controllers.CreateStreamController)
		stream.POST("/stop-stream", controllers.StopActiveStream)
		stream.GET("/all-active-stream", controllers.GetAllActiveStreamController)
		stream.GET("/by-stream-id", controllers.GetActiveStreamByStreamIDController)
		stream.GET("/by-streamer-id", controllers.GetActiveStreamByStreamerIDController)
	}
}
