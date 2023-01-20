package ducksms

// CancelSmsSchedule struct for CancelSmsSchedule
type CancelSmsSchedule struct {
	Status  int32                             `json:"status,omitempty"`
	Message string                            `json:"message,omitempty"`
	Data    map[string]map[string]interface{} `json:"data,omitempty"`
}
