package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(api *gin.RouterGroup) {
	category := api.Group("/category")
	{
		category.POST("/create-category", controllers.CreateCategoryController)
		category.GET("/get-all-category", controllers.GetAllCategoryController)
	}
}
