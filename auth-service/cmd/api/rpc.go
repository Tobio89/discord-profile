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

func (r *RPCServer) GetLoginURL(payload RPCPayload, resp *string) error {

	log.Println("auth: received request from: ", payload.User)

	*resp = "http://localhost:5173/" + payload.User

	return nil
}

func (r *RPCServer) RequestLogin(payload RPCLoginPayload, resp *string) error {
	log.Println("auth: received login request for user: ", payload.Username)

	*resp = "http://localhost:5173/" + "?user=" + payload.Username + "&id=" + payload.ID

	return nil

}
