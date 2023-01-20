package ducksms

// ErrorUnauthenticated struct for ErrorUnauthenticated
type ErrorUnauthenticated struct {
	Status  int32  `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
