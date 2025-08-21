package dtos

type ModerationRequest struct {
	Text string `json:"text"`
}

type ModerationResponse struct {
	IsInappropriate bool `json:"is_inappropriate"`
}
