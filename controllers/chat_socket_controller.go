package controllers

import (
	"log"
	"main/config"
	"main/dtos"
	"main/services"

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

	for {
		var incoming dtos.SocketMessage
		err := conn.ReadJSON(&incoming)
		if err != nil {
			log.Println("Read error: ", err)
			delete(chatRooms[streamID], conn)
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
		log.Println("Decode error: ", err)
		return
	}
	chatMsg.StreamID = streamID
	broadcast <- chatMsg
}

func handleMessages() {
	for {
		msg := <-broadcast

		if err := services.SaveMessage(msg.StreamID, msg.UserID, msg.Content); err != nil {
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
