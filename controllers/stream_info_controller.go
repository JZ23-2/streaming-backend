package controllers

import (
	"log"
	"main/dtos"
	"main/services"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateStreamInfo godoc
// @Summary      Create or update stream info
// @Description  Create or update stream info
// @Tags         StreamInfo
// @Accept       application/json
// @Produce      json
// @Param        dto body  dtos.CreateStreamInfoRequest true "DTO"
// @Success      204 {string} string "No Content"
// @Failure      400 {object} map[string]string
// @Failure      409 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /stream-info/ [put]
func CreateOrUpdateStreamInfo(c *gin.Context) {
	var dto dtos.CreateStreamInfoRequest

	if err := c.ShouldBindJSON(&dto); err != nil {
		utils.FailedResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	if err := services.CreateOrUpdateStreamInfo(dto); err != nil {
		utils.FailedResponse(c, http.StatusInternalServerError, "an error occurred while saving stream info")
		return
	}

	utils.SuccessResponse(c, http.StatusNoContent, "stream info saved successfully", nil)
}

// GetStreamInfoByUserID godoc
// @Summary      get stream info by user id
// @Description  get stream info by user id
// @Tags         StreamInfo
// @Accept       application/json
// @Produce      json
// @Param        hostPrincipalID path string true "Host Principal ID"
// @Success      200 {object} dtos.StreamInfoResponse
// @Failure      400 {object} map[string]string
// @Failure      409 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /stream-info/{hostPrincipalID} [get]
func GetStreamInfoByUserID(c *gin.Context) {
	hostPrincipalID := c.Param("hostPrincipalID")
	log.Println("sdfaf", hostPrincipalID)
	streamInfo, err := services.GetStreamInfoByUserID(hostPrincipalID)
	if err != nil {
		utils.FailedResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	resp := dtos.StreamInfoResponse{
		HostPrincipalID: streamInfo.HostPrincipalID,
		Title:           streamInfo.Title,
		CategoryName:    streamInfo.Category.CategoryName,
	}

	utils.SuccessResponse(c, http.StatusOK, "stream info retrieved", resp)

}
