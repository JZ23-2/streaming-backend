package dtos

type SocketMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
