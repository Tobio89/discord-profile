package rpc

import (
	"log"
	"net/rpc"
)

type RPCServer struct{}

type RPCPayload struct {
	User string
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {

	log.Println("broker: received request from: ", payload.User)

	*resp = "received login request from: " + payload.User

	return nil
}

func MakeRPCCall(username string) {
	client, err := rpc.Dial("tcp", "localhost:5001")
	if err != nil {
		log.Println("while creating RPC client: ", err)
		return
	}

	rpcPayload := RPCPayload{
		User: username,
	}

	// the response from the RPC call
	var result string

	err = client.Call("RPCServer.LogInfo", rpcPayload, &result)
	if err != nil {
		log.Println("while calling RPC: ", err)
		return
	}

	log.Println("response from RPC: ", result)

}
