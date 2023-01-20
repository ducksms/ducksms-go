package ducksms

// SmsSchema struct for SmsSchema
type SmsSchema struct {
	// Preview the sms information
	Preview       string    `json:"preview,omitempty"`
	MobileNumbers *[]string `json:"mobile_numbers,omitempty"`
	Message       *string   `json:"message,omitempty"`
	SenderId      *string   `json:"sender_id,omitempty"`
	ScheduledAt   *string   `json:"scheduled_at,omitempty"`
	CallbackUrl   *string   `json:"callback_url,omitempty"`
}
