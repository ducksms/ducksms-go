package ducksms

// ErrorException struct for ErrorException
type ErrorException struct {
	Status  int32  `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
