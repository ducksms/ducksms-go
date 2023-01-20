package ducksms

// ErrorValidation struct for ErrorValidation
type ErrorValidation struct {
	Status  int32               `json:"status,omitempty"`
	Message string              `json:"message,omitempty"`
	Errors  map[string][]string `json:"errors,omitempty"`
}
