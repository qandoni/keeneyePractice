package enum

type EmailStatus string

const (
	StatusEmailPending EmailStatus = "pending"
	StatusEmailSent    EmailStatus = "sent"
	StatusEmailFailed  EmailStatus = "failed"
	StatusEmailGiveUp  EmailStatus = "give up"
	StatusEmailSending EmailStatus = "sending"
)
