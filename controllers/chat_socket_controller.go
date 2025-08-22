package controllers

import (
	"log"
	"main/config"
	"main/dtos"
	"main/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

var chatRooms = make(map[string]map[*websocket.Conn]bool)
var broadcast = make(chan dtos.ChatMessage)

func init() {
	go handleMessages()
}

func HandleWebSocket(c *gin.Context) {
	streamID := c.Param("streamID")
	conn, err := config.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Upgrade error: ", err)
		return
	}

	defer conn.Close()

	if chatRooms[streamID] == nil {
		chatRooms[streamID] = make(map[*websocket.Conn]bool)
	}

	chatRooms[streamID][conn] = true

	broadcastViewerCount(streamID)

	for {
		var incoming dtos.SocketMessage
		err := conn.ReadJSON(&incoming)
		if err != nil {
			log.Println("Read error: ", err)
			delete(chatRooms[streamID], conn)
			broadcastViewerCount(streamID)
			break
		}

		switch incoming.Type {
		case "chat_message":
			handleChatMessage(streamID, incoming.Data)
		}
	}
}

func handleChatMessage(streamID string, data interface{}) {
	var chatMsg dtos.ChatMessage

	if err := mapstructure.Decode(data, &chatMsg); err != nil {
		log.Println("Decode error:", err)
		return
	}

	filteredText, _ := services.ModerateMessage(chatMsg.Content)

	chatMsg.Content = filteredText
	chatMsg.StreamID = streamID
	broadcast <- chatMsg
}

func handleMessages() {
	for {
		msg := <-broadcast
		log.Println("saving message", msg)
		if err := services.SaveMessage(msg.StreamID, msg.UserID, msg.Username, msg.Content); err != nil {
			log.Println("Error saving message: ", err)
		}

		for conn := range chatRooms[msg.StreamID] {
			err := conn.WriteJSON(map[string]interface{}{
				"type": "chat_message",
				"data": msg,
			})
			if err != nil {
				log.Println("Write error: ", err)
				conn.Close()
				delete(chatRooms[msg.StreamID], conn)
			}
		}
	}
}

func broadcastViewerCount(streamID string) {
	count := len(chatRooms[streamID])
	msg := dtos.SocketMessage{
		Type: "viewer_count",
		Data: count,
	}

	for conn := range chatRooms[streamID] {
		if err := conn.WriteJSON(msg); err != nil {
			log.Println("Write error:", err)
			conn.Close()
			delete(chatRooms[streamID], conn)
		}
	}
}

func HandleGetViewerCount(c *gin.Context) {
	streamID := c.Param("streamID")
	count := GetViewerCount((streamID))

	c.JSON(http.StatusOK, count)
}

func GetViewerCount(streamID string) int {
	if conns, ok := chatRooms[streamID]; ok {
		return len(conns)
	}

	return 0
}
