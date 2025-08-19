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
	file, err := c.FormFile("thumbnail")
	if err != nil {
		utils.FailedResponse(c, http.StatusBadRequest, "thumbnail file is required")
		return
	}
	req.Thumbnail = file

	resp, err := services.CreateStream(req)
	if err != nil {
		log.Println(err)
		utils.FailedResponse(c, http.StatusConflict, "failed to create stream ")
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
// @Success      200   {object}  []dtos.StreamResponse
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
// @Success      200   {object}  dtos.StreamResponse
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

// StopActiveStream godoc
// @Summary      Stop active stream
// @Description  Stop active stream
// @Tags         Stream
// @Accept       json
// @Produce      json
// @Param        dto body dtos.UpdateStreamActiveStatusRequest true "DTO"
// @Success      200   {object}  dtos.StreamResponse
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /streams/stop-stream [post]
func StopActiveStream(c *gin.Context) {
	var req dtos.UpdateStreamActiveStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.FailedResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	resp, err := services.StopStream(req.HostPrincipalID)

	if err != nil {
		utils.FailedResponse(c, http.StatusInternalServerError, "an error occurred while updating stream status")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "update success", resp)
}

// GetActiveStreamByStreamerID godoc
// @Summary      Get active stream by streamerID
// @Description  Get active stream by streamerID
// @Tags         Stream
// @Accept       json
// @Produce      json
// @Param        streamerID query string true "Streamer ID"
// @Success      200   {object}  dtos.StreamResponse
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /streams/by-streamer-id [get]
func GetActiveStreamByStreamerIDController(c *gin.Context) {
	streamerID := c.Query("streamerID")

	if streamerID == "" {
		utils.FailedResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	stream, err := services.GetActiveStreamByStreamerID(streamerID)
	if err != nil {
		utils.FailedResponse(c, http.StatusInternalServerError, "something wrong")
		return
	}

	resp := dtos.StreamResponse{
		StreamID:        stream.StreamID,
		HostPrincipalID: stream.HostPrincipalID,
		ThumbnailURL:    stream.ThumbnailURL,
		IsActive:        stream.IsActive,
		CreatedAt:       stream.CreatedAt,
	}

	if stream.StreamInfoID != nil {
		resp.Title = stream.StreamInfo.Title
		if stream.StreamInfo.StreamCategoryID != nil {
			resp.CategoryName = stream.StreamInfo.Category.CategoryName
		}
	}

	utils.SuccessResponse(c, 200, "success", resp)
}
