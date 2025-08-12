package controllers

import (
	"main/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMessagesByStreamID godoc
// @Summary      Get all messages for a stream
// @Description  Retrieve all chat messages by Stream ID
// @Tags         messages
// @Param        stream_id   path      string  true  "Stream ID"
// @Produce      json
// @Success      200  {array}   models.Message
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /messages/{stream_id} [get]
func GetMessagesByStreamID(c *gin.Context) {
	streamID := c.Param("stream_id")

	if streamID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "stream_id is required"})
		return
	}

	messages, err := services.GetMessagesByStreamID(streamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}
