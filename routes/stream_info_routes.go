package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func StreamInfoRoutes(api *gin.RouterGroup) {
	streamInfo := api.Group("/stream-info")
	{
		streamInfo.PUT("", controllers.CreateOrUpdateStreamInfo)
		streamInfo.GET("/:hostPrincipalID", controllers.GetStreamInfoByUserID)

	}
}
