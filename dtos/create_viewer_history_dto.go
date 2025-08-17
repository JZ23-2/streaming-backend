package dtos

type CreateViewerHistoryRequest struct {
	ViewerHistoryPrincipalID string `json:"viewerHistoryPrincipalID" example:"user123"`
	ViewerHistoryStreamID    string `json:"viewerHistoryStreamID" example:"stream123"`
}

type CreateViewerHistoryResponse struct {
	ViewerHistoryID          string `json:"viewerHistoryID"`
	ViewerHistoryPrincipalID string `json:"viewerHistoryPrincipalID"`
	ViewerHistoryStreamID    string `json:"viewerHistoryStreamID"`
}
