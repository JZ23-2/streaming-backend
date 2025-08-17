package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func ViewerHistoryRoutes(api *gin.RouterGroup) {
	category := api.Group("/viewer-history")
	{
		category.POST("/create", controllers.CreateViewerHistoryController)
	}
}
