package ducksms

// CreditHistory struct for CreditHistory
type CreditHistory struct {
	Status  int32              `json:"status,omitempty"`
	Message string             `json:"message,omitempty"`
	Data    map[string][]int32 `json:"data,omitempty"`
}
