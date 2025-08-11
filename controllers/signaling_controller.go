package controllers

import (
	"main/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pion/webrtc/v3"
)

func HandlePublish(c *gin.Context) {
	var offer webrtc.SessionDescription

	if err := c.ShouldBindJSON(&offer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid SDP offer"})
		return
	}

	_, answer, err := services.CreatePublisherPC(offer)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, answer)
}

func HandleView(c *gin.Context) {
	var offer webrtc.SessionDescription

	if err := c.ShouldBindJSON(&offer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SDP offer"})
		return
	}

	_, answer, _, err := services.CreateViewerPC(offer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, answer)
}
