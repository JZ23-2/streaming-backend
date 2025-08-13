package controllers

import (
	"log"
	"main/config"
	"main/dtos"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

var globalRooms = make(map[string]map[*websocket.Conn]bool)

func HandleGlobalSocket(c *gin.Context) {
	principalID := c.Param("principalID")
	conn, err := config.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Upgrade error: ", err)
		return
	}

	defer conn.Close()

	if globalRooms[principalID] == nil {
		globalRooms[principalID] = make(map[*websocket.Conn]bool)
	}

	globalRooms[principalID][conn] = true

	for {
		var incoming dtos.SocketMessage
		err := conn.ReadJSON(&incoming)
		if err != nil {
			log.Println("Read error: ", err)
			delete(globalRooms[principalID], conn)
			break
		}

		switch incoming.Type {
		case "start_stream":
			handleStartStream(incoming.Data)
		}
	}
}

func handleStartStream(data interface{}) {
	var streamMsg dtos.StartStreamMessage
	if err := mapstructure.Decode(data, &streamMsg); err != nil {
		log.Println("Decode error: ", err)
		return
	}

	for _, followerID := range streamMsg.Followers {
		conns := globalRooms[followerID]
		for conn := range conns {
			err := conn.WriteJSON(map[string]interface{}{
				"type": "stream_started",
				"data": map[string]string{
					"stream_id":   streamMsg.StreamID,
					"streamer_id": streamMsg.StreamerID,
				},
			})

			if err != nil {
				log.Println("Write error: ", err)
				conn.Close()
				delete(conns, conn)
			}
		}
	}
}
