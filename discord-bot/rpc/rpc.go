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

type RPCSignupPayload struct {
	Username string
	ID       string
	Token    string
}

type RPCSignupResponse struct {
	AlreadyExists bool
	Message       string
}

func DialRPCServer() (*rpc.Client, error) {
	client, err := rpc.Dial("tcp", "profile-broker:5001")
	if err != nil {
		return nil, err
	}

	return client, nil

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

func RPCRequestSignup(signupReq RPCSignupPayload) (result RPCSignupResponse, err error) {
	client, err := DialRPCServer()
	if err != nil {
		log.Println("while dialing RPC server: ", err)
		return RPCSignupResponse{}, err
	}

	payload := RPCSignupPayload{
		Username: signupReq.Username,
		ID:       signupReq.ID,
		Token:    signupReq.Token,
	}

	err = client.Call("RPCServer.RequestSignup", payload, &result)
	if err != nil {
		log.Println("while calling RPC: ", err)
		return RPCSignupResponse{}, err
	}

	log.Println("response from RPC: ", result)

	return result, nil
}
