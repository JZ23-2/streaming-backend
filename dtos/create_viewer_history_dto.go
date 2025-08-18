package dtos

type CreateViewerHistoryRequest struct {
	ViewerHistoryPrincipalID     string `json:"viewerHistoryPrincipalID" example:"user123"`
	ViewerHistoryStreamHistoryID string `json:"viewerHistoryStreamHistoryID" example:"stream123"`
}

type CreateViewerHistoryResponse struct {
	ViewerHistoryID              string `json:"viewerHistoryID"`
	ViewerHistoryPrincipalID     string `json:"viewerHistoryPrincipalID"`
	ViewerHistoryStreamHistoryID string `json:"viewerHistoryStreamHistoryID"`
}
