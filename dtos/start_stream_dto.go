package dtos

type StartStreamMessage struct {
	StreamerID string   `json:"streamerId"`
	StreamID   string   `json:"streamId"`
	Followers  []string `json:"followers"`
}
