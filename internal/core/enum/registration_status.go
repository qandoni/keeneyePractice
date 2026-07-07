package enum

type RegistrationStatus string

const (
	StatusPending   RegistrationStatus = "pending"
	StatusCompleted RegistrationStatus = "completed"
	StatusExpired   RegistrationStatus = "expired"
)
