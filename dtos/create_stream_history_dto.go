package dtos

type CreateStreamHistoryRequest struct {
	StreamHistoryStreamID string `json:"streamHistoryStreamId" example:"stream123"`
	HostPrincipalID       string `json:"hostPrincipalId" example:"user123"`
	VideoUrl              string `json:"videoUrl" example:"supabase storage video"`
}

type CreateStreamHistoryResponse struct {
	StreamHistoryID       string `json:"streamHistoryId"`
	StreamHistoryStreamID string `json:"streamHistoryStreamId"`
	HostPrincipalID       string `json:"hostPrincipalId"`
	VideoUrl              string `json:"videoUrl"`
	Duration              int    `json:"duration"`
}
