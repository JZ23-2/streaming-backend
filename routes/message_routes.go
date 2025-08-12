package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func MessageRoutes(api *gin.RouterGroup) {
	message := api.Group("/messages")
	{
		message.GET("/:stream_id", controllers.GetMessagesByStreamID)
	}
}
