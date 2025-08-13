package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func GlobalSocketRoutes(api *gin.RouterGroup) {
	globalSocket := api.Group("/global-sockets")
	{
		globalSocket.GET("/ws/:principalID", controllers.HandleGlobalSocket)
	}
}
