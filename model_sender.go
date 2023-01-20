package ducksms

// Sender struct for Sender
type Sender struct {
	// Sender name
	Name string `json:"name,omitempty"`
	// Sender description
	Description string `json:"description,omitempty"`
	// Default sender id
	Default *bool `json:"default,omitempty"`
	// Sender id status
	Status string `json:"status,omitempty"`
}
