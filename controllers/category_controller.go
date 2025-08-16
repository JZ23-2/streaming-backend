package controllers

import (
	"main/dtos"
	"main/services"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCategory godoc
// @Summary      Create Category
// @Description  Create Category
// @Tags         Category
// @Accept       json
// @Produce      json
// @Param        category  body      dtos.CreateCategoryRequest  true  "Category Name"
// @Success      201   {object}  models.Category
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /category/create-category [post]
func CreateCategoryController(c *gin.Context) {
	var req dtos.CreateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.FailedResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	resp, err := services.CreateCategory(req)

	if err != nil {
		utils.FailedResponse(c, http.StatusConflict, "failed to create category")
		return
	}

	utils.SuccessResponse(c, 201, "category created", resp)
}

// GetAllCategory godoc
// @Summary      Get all categories
// @Description  Retrieve all categories from the database
// @Tags         Category
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Category
// @Failure      404  {object}  map[string]string
// @Router       /category/get-all-category [get]
func GetAllCategoryController(c *gin.Context) {

	resp, err := services.GetAllCategory()

	if err != nil {
		utils.FailedResponse(c, http.StatusNotFound, "failed to retrieve category data")
		return
	}

	utils.SuccessResponse(c, 200, "category retrieved", resp)
}
