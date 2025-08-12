package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func ChatRoutes(api *gin.RouterGroup) {
	chat := api.Group("/chats")
	{
		chat.GET("/ws/:streamID", controllers.HandleWebSocket)
	}
}
