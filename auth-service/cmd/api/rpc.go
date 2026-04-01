package main

import "log"

type RPCServer struct{}

type RPCPayload struct {
	User string
}

type RPCLoginPayload struct {
	Username string
	ID       string
	Token    string
}

type RPCLoginResponse struct {
	Success bool
	URL     string
	Message string
}

type RPCSignupPayload struct {
	Username string
	ID       string
	Token    string
}

type RPCSignupResponse struct {
	AlreadyExists bool
	Message       string
}

func (r *RPCServer) GetLoginURL(payload RPCPayload, resp *string) error {

	log.Println("auth: received request from: ", payload.User)

	*resp = "http://localhost:5173/" + payload.User

	return nil
}

func (r *RPCServer) RequestLogin(payload RPCLoginPayload, resp *string) error {
	return app.HandleLoginRequest(payload, resp)
}

func (r *RPCServer) RequestSignup(payload RPCSignupPayload, resp *RPCSignupResponse) error {
	return app.HandleSignupRequest(payload, resp)
}
