package models

type (
	KeigoRequest struct {
		Body string `json:"body" binding:"required,omitempty"`
	}
	KeigoResponse struct {
		ConvertedBody string `json:"converted_body" binding:"required,omitempty"`
	}
)
