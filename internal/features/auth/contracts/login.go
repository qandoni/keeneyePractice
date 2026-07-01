package auth_contracts

type LoginInput struct {
	Login    string
	Password string
}

type LoginOutput struct {
	AccessToken  string
	RefreshToken string
}
