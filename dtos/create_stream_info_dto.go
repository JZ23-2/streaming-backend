package dtos

type CreateStreamInfoRequest struct {
	HostPrincipalID string `json:"hostPrincipalId" example:"123"`
	Title           string `json:"title" example:"My First Stream!!"`
	CategoryID      string `json:"categoryId" example:"123"`
}
