package main

import "log"

type RPCServer struct{}

type RPCPayload struct {
	User string
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {

	log.Println("broker: received request from: ", payload.User)

	*resp = "received login request from: " + payload.User

	return nil
}
