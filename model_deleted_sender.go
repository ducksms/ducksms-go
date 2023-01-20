package ducksms

// DeletedSender struct for DeletedSender
type DeletedSender struct {
	Status  int32                             `json:"status,omitempty"`
	Message string                            `json:"message,omitempty"`
	Data    map[string]map[string]interface{} `json:"data,omitempty"`
}
