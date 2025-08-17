package controllers

import (
	"log"
	"main/dtos"
	"main/services"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateStreamHistory godoc
// @Summary      Create Stream History
// @Description  Create Stream History
// @Tags         Stream History
// @Accept       json
// @Produce      json
// @Param        category  body      dtos.CreateStreamHistoryRequest  true  "Stream History"
// @Success      201   {object}  dtos.CreateStreamHistoryResponse
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /stream-history/create [post]
func CreateStreamHistoryController(c *gin.Context) {

	var req dtos.CreateStreamHistoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.FailedResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	var resp, err = services.CreateStreamHistory(req)

	if err != nil {
		log.Println(err)
		utils.FailedResponse(c, http.StatusConflict, "failed to create")
		return
	}

	utils.SuccessResponse(c, 201, "stream history success", resp)
}

// GetAllStreamHistoryByStreamerID godoc
// @Summary      Get all stream history by streamerID
// @Description  Get all stream history by streamerID
// @Tags         Stream History
// @Accept       json
// @Produce      json
//
//	@Param			hostPrincipalID	query		string	true	"Host Principal ID"
//
// @Success      201   {object}  []dtos.GetAllStreamResponse
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /stream-history/all-stream [get]
func GetAllStreamHistoryByStreamerIDController(c *gin.Context) {
	req := c.Query("hostPrincipalID")

	resp, err := services.GetAllStreamHistoryByStreamerID(req)

	if err != nil {
		utils.FailedResponse(c, http.StatusNotFound, "history not found")
		return
	}

	utils.SuccessResponse(c, 200, "stream history retrieved!", resp)
}

// GetStreamHistoryByStreamHistoryID godoc
// @Summary      Get stream history by streamHistoryID
// @Description  Get stream history by streamHistoryID
// @Tags         Stream History
// @Accept       json
// @Produce      json
//
//	@Param			streamHistoryID	query		string	true	"Stream History ID"
//
// @Success      201   {object}  dtos.GetAllStreamResponse
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /stream-history/by-id [get]
func GetAllStreamHistoryByIdController(c *gin.Context) {
	req := c.Query("streamHistoryID")
	if req == "" {
		utils.FailedResponse(c, http.StatusBadRequest, "streamHistoryID is required")
		return
	}

	resp, err := services.GetAllStreamHistoryByID(req)

	if err != nil {
		log.Println(err)
		utils.FailedResponse(c, http.StatusNotFound, "history not found")
		return
	}

	utils.SuccessResponse(c, 200, "stream history retrieved!", resp)
}
