package dtos

type CreateHighlightRequest struct {
	Clips []ClipRequest `json:"clips"`
}

type ClipRequest struct {
	HighlightStreamHistoryID string `json:"highlightStreamHistoryID" example:"streamhistory123"`
	HighlightUrl             string `json:"highlightUrl" example:"url123"`
	StartHighlight           string `json:"startHighlight" example:"0.0"`
	EndHighlight             string `json:"endHighlight" example:"0.5"`
	HighlightDescription     string `json:"highlightDescription" example:"hello world"`
}

type CreateHighlightResponse struct {
	Highlights []ClipResponse `json:"highlights"`
}

type ClipResponse struct {
	HighlightID              string `json:"highlightID"`
	HighlightStreamHistoryID string `json:"highlightStreamHistoryID"`
	HighlightUrl             string `json:"highlightUrl"`
	StartHighlight           string `json:"startHighlight"`
	EndHighlight             string `json:"endHighlight"`
	HighlightDescription     string `json:"highlightDescription"`
}
