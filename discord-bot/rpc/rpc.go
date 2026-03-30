package rpc

import (
	"log"
	"net/rpc"
)

type RPCPayload struct {
	User string
}

type RPCLoginPayload struct {
	Username string
	ID       string
	Token    string
}
type LoginRequest struct {
	Username string
	ID       string
	Token    string
}

func DialRPCServer() (*rpc.Client, error) {
	client, err := rpc.Dial("tcp", "profile-broker:5001")
	if err != nil {
		return nil, err
	}

	return client, nil

}

func RPCRequestLogin(loginReq LoginRequest) (result string, err error) {
	client, err := DialRPCServer()
	if err != nil {
		log.Println("while dialing RPC server: ", err)
		return "", err
	}

	payload := RPCLoginPayload{
		Username: loginReq.Username,
		ID:       loginReq.ID,
		Token:    loginReq.Token,
	}

	err = client.Call("RPCServer.RequestLogin", payload, &result)
	if err != nil {
		log.Println("while calling RPC: ", err)
		return "", err
	}

	log.Println("response from RPC: ", result)

	return result, nil
}

func MakeRPCCall(username string) (result string) {
	client, err := DialRPCServer()
	if err != nil {
		log.Println("while dialing RPC server: ", err)
		return
	}

	rpcPayload := RPCPayload{
		User: username,
	}

	err = client.Call("RPCServer.LogInfo", rpcPayload, &result)
	if err != nil {
		log.Println("while calling RPC: ", err)
		return
	}

	log.Println("response from RPC: ", result)

	return result

}
