package controllers

import (
	"main/dtos"
	"main/services"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateViewerHistory godoc
// @Summary      Create Viewer History
// @Description  Create Viewer History
// @Tags         Viewer History
// @Accept       json
// @Produce      json
// @Param        category  body      dtos.CreateViewerHistoryRequest  true  "Viewer History"
// @Success      201   {object}  dtos.CreateViewerHistoryResponse
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /viewer-history/create [post]
func CreateViewerHistoryController(c *gin.Context) {

	var req dtos.CreateViewerHistoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.FailedResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	resp, err := services.CreateViewerHistory(req)

	if err != nil {
		utils.FailedResponse(c, http.StatusConflict, "failed to create viewer history")
		return
	}

	utils.SuccessResponse(c, 201, "viewer history create success", resp)
}
