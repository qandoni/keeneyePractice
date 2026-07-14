package kafka

type RegistrationEmailMessage struct {
	RequestID int `json:"request_id"`
}

type RetryMessage struct {
	RequestID int `json:"request_id"`
	Retry     int `json:"retry"`
}
