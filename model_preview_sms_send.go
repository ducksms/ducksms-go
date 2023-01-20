package ducksms

// PreviewSmsSend struct for PreviewSmsSend
type PreviewSmsSend struct {
	Status  int32               `json:"status,omitempty"`
	Message string              `json:"message,omitempty"`
	Data    map[string][]string `json:"data,omitempty"`
}
