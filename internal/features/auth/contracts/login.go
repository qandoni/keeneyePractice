package auth_contracts

type LoginInput struct {
	Login    string
	Password string
}

type LoginOuput struct {
	Token string
}
