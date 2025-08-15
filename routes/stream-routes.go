package routes

import "github.com/gin-gonic/gin"

func StreamRoutes(api *gin.RouterGroup) {
	stream := api.Group("streams")
	{
		stream.POST("/start")
	}
}
