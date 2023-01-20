package ducksms

// ErrorInvalidRequest struct for ErrorInvalidRequest
type ErrorInvalidRequest struct {
	Status  int32  `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
