package controllers

import (
	"main/dtos"
	"main/services"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateVideoHightlight godoc
// @Summary      Create Video Highlight
// @Description  Create Video Highlight
// @Tags         Video Highlight
// @Accept       json
// @Produce      json
// @Param        highlight  body      dtos.CreateHighlightRequest  true  "highlight"
// @Success      201   {object}  dtos.CreateHighlightResponse
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /highlight/create [post]
func CreateHighlightController(c *gin.Context) {
	var req dtos.CreateHighlightRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.FailedResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	resp, err := services.CreateHighlight(req)

	if err != nil {
		utils.FailedResponse(c, http.StatusInternalServerError, "failed to create highlight")
		return
	}

	utils.SuccessResponse(c, 201, "highlight created", resp)
}

// GetHighlightByStreamerID godoc
// @Summary      Get Highlight by StreamerID
// @Description  Get Highlight by StreamerID
// @Tags         Video Highlight
// @Accept       json
// @Produce      json
// @Param        streamerID query string true "Streamer ID"
// @Success      200   {object}  dtos.CreateHighlightResponse
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /highlight/by-streamer-id [get]
func GetAllHighlightByStreamerIDController(c *gin.Context) {
	streamerID := c.Query("streamerID")

	if streamerID == "" {
		utils.FailedResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	resp, err := services.GetAllHighlightByStreamerID(streamerID)

	if err != nil {
		utils.FailedResponse(c, http.StatusInternalServerError, "failed to retrieve highlight")
		return
	}

	utils.SuccessResponse(c, 200, "retrieved highlight success", resp)
}
