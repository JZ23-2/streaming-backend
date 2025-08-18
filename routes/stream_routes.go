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
		stream.GET("/all-active-stream", controllers.GetAllActiveStreamController)
		stream.GET("/by-stream-id", controllers.GetActiveStreamByStreamIDController)
		stream.PATCH("/update-active-status", controllers.UpdateStreamActiveStatusController)
		stream.PATCH("/update-stream", controllers.UpdateStreamController)
	}
}
