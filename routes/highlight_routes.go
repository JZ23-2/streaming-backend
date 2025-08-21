package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func Highlight(api *gin.RouterGroup) {
	Highlight := api.Group("/highlight")
	{
		Highlight.POST("/create", controllers.CreateHighlightController)
	}
}
