package ducksms

// CreditBalance struct for CreditBalance
type CreditBalance struct {
	Status  int32              `json:"status,omitempty"`
	Message string             `json:"message,omitempty"`
	Data    map[string][]int32 `json:"data,omitempty"`
}
