package registration_contracts

type CompleteInput struct {
	Token    string
	Password string
}

type CompleteOutput struct {
	UserID int
}
