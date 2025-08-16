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
