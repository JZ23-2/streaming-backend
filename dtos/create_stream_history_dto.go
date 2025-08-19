package dtos

type CreateStreamHistoryRequest struct {
	HostPrincipalID string `json:"hostPrincipalID" example:"user123"`
	VideoUrl        string `json:"videoUrl" example:"supabase storage video"`
}

type CreateStreamHistoryResponse struct {
	StreamHistoryID       string `json:"streamHistoryId"`
	StreamHistoryStreamID string `json:"streamHistoryStreamId"`
	HostPrincipalID       string `json:"hostPrincipalId"`
	VideoUrl              string `json:"videoUrl"`
	Duration              int    `json:"duration"`
}
