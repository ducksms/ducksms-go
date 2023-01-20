package ducksms

// GetSender struct for GetSender
type GetSender struct {
	Status  int32              `json:"status,omitempty"`
	Message string             `json:"message,omitempty"`
	Data    map[string][]int32 `json:"data,omitempty"`
}
