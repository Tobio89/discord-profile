package main

import "log"

type RPCServer struct{}

type RPCPayload struct {
	User string
}

func (r *RPCServer) GetLoginURL(payload RPCPayload, resp *string) error {

	log.Println("auth: received request from: ", payload.User)

	*resp = "http://google.com/"

	return nil
}
