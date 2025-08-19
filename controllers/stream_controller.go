package controllers

import (
	"log"
	"main/dtos"
	"main/services"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateStream godoc
// @Summary      Create Stream
// @Description  Create Stream with thumbnail upload
// @Tags         Stream
// @Accept       multipart/form-data
// @Produce      json
// @Param        hostPrincipalId  formData  string true  "Host Principal ID"
// @Param        title            formData  string true  "Stream title"
// @Param        streamCategoryId formData  string true  "Stream Category ID"
// @Param        thumbnail        formData  file   true  "Thumbnail file"
// @Success      201 {object} dtos.CreateStreamingResponse
// @Failure      400 {object} map[string]string
// @Failure      409 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /streams/create-stream [post]
func CreateStreamController(c *gin.Context) {
	var req dtos.CreateStreamingRequest

	req.HostPrincipalID = c.PostForm("hostPrincipalId")
	req.Title = c.PostForm("title")
	req.StreamCategoryID = c.PostForm("streamCategoryId")

	file, err := c.FormFile("thumbnail")
	if err != nil {
		utils.FailedResponse(c, http.StatusBadRequest, "thumbnail file is required")
		return
	}
	req.Thumbnail = file

	resp, err := services.CreateStream(req)
	if err != nil {
		log.Println(err)
		utils.FailedResponse(c, http.StatusConflict, "failed to create stream: ")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "stream created", resp)
}

// GetAllActiveStream godoc
// @Summary      Get all active stream
// @Description  Get all active stream
// @Tags         Stream
// @Accept       json
// @Produce      json
// @Success      200   {object}  []dtos.GetActiveAllStreamResponse
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /streams/all-active-stream [get]
func GetAllActiveStreamController(c *gin.Context) {

	resp, err := services.GetAllActiveStream()

	if err != nil {
		utils.FailedResponse(c, 404, "active stream not found!")
		return
	}

	for i := range resp {
		resp[i].ViewerCount = GetViewerCount(resp[i].HostPrincipalID)
	}

	utils.SuccessResponse(c, 200, "active stream found!", resp)
}

// GetActiveStreamByStreamID godoc
// @Summary      Get active stream by streamID
// @Description  Get active stream by streamID
// @Tags         Stream
// @Accept       json
// @Produce      json
// @Param        streamID query string true "Stream ID"
// @Success      200   {object}  dtos.GetActiveAllStreamResponse
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /streams/by-stream-id [get]
func GetActiveStreamByStreamIDController(c *gin.Context) {
	streamID := c.Query("streamID")

	if streamID == "" {
		utils.FailedResponse(c, http.StatusBadRequest, "Invalid Request")
		return
	}

	resp, err := services.GetActiveStreamByStreamID(streamID)

	if err != nil {
		utils.FailedResponse(c, 404, "active stream not found")
		return
	}

	utils.SuccessResponse(c, 200, "active stream found", resp)
}

// UpdateStreamActiveStatus godoc
// @Summary      Update stream active status
// @Description  Update stream active status
// @Tags         Stream
// @Accept       json
// @Produce      json
// @Param        UpdateStream body dtos.UpdateStreamActiveStatusRequest true "Update Stream"
// @Success      200   {object}  dtos.GetActiveAllStreamResponse
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /streams/update-active-status [patch]
func UpdateStreamActiveStatusController(c *gin.Context) {
	var req dtos.UpdateStreamActiveStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.FailedResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	resp, err := services.UpdateStreamActiveStatus(req)

	if err != nil {
		utils.FailedResponse(c, http.StatusConflict, "update went wrong")
		return
	}

	utils.SuccessResponse(c, 200, "update success", resp)
}

// UpdateStream godoc
// @Summary      Update Stream
// @Description  Update Stream
// @Tags         Stream
// @Accept       multipart/form-data
// @Produce      json
// @Param        streamId  		  formData  string true  "Stream ID"
// @Param        title            formData  string true  "Stream title"
// @Param        streamCategoryId formData  string true  "Stream Category ID"
// @Param        thumbnail        formData  file   false  "Thumbnail file"
// @Success      200 {object} dtos.UpdateStreamingResponse
// @Failure      400 {object} map[string]string
// @Failure      409 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /streams/update-stream [patch]
func UpdateStreamController(c *gin.Context) {
	var req dtos.UpdateStreamingRequest

	req.StreamID = c.PostForm("streamId")
	req.Title = c.PostForm("title")
	req.StreamCategoryID = c.PostForm("streamCategoryId")

	file, err := c.FormFile("thumbnail")
	if err == nil {
		req.Thumbnail = file
	} else {
		req.Thumbnail = nil
	}

	resp, err := services.UpdateStream(req)
	if err != nil {
		log.Println(err)
		utils.FailedResponse(c, http.StatusConflict, "failed to update stream")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "stream updated", resp)
}
