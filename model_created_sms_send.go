package ducksms

// CreatedSmsSend struct for CreatedSmsSend
type CreatedSmsSend struct {
	Status  int32               `json:"status,omitempty"`
	Message string              `json:"message,omitempty"`
	Data    map[string][]string `json:"data,omitempty"`
}
