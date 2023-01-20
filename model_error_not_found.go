package ducksms

// ErrorNotFound struct for ErrorNotFound
type ErrorNotFound struct {
	Status  int32  `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
