package rpccontracts

type Payload struct {
	User string
}

type LoginPayload struct {
	Username string
	ID       string
	Token    string
}

type LoginResponse struct {
	Success bool
	Error   bool
	URL     string
	Message string
}

type SignupPayload struct {
	Username string
	ID       string
	Token    string
}

type SignupResponse struct {
	AlreadyExists bool
	Message       string
}

type TokenCheckPayload struct {
	Token string
}

type TokenCheckResponse struct {
	UserID  string
	Message string
	JWT     string
}
