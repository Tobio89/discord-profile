package main

import (
	"log"
	"net/rpc"
)

type RPCServer struct{}

type RPCPayload struct {
	User string
}

type RPCLoginPayload struct {
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

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {

	log.Println("broker: received request from: ", payload.User)

	response := MakeRPCCall(payload.User)

	*resp = response

	return nil
}

func (r *RPCServer) RequestLogin(payload RPCLoginPayload, resp *string) error {
	log.Println("broker: received login request for user: ", payload.Username)

	response, err := RPCRequestLogin(payload)
	if err != nil {
		log.Println("error while requesting login: ", err)
	}

	*resp = response

	return nil

}

func (r *RPCServer) RequestSignup(payload RPCSignupPayload, resp *RPCSignupResponse) error {
	log.Println("broker: received signup request for user: ", payload.Username)

	response, err := RPCRequestSignup(payload)
	if err != nil {
		log.Println("error while requesting signup: ", err)
		return err
	}

	*resp = response

	return nil

}

func DialRPCServer() (*rpc.Client, error) {
	client, err := rpc.Dial("tcp", "auth-service:5001")
	if err != nil {
		return nil, err
	}

	return client, nil

}

func RPCRequestLogin(loginReq RPCLoginPayload) (result string, err error) {
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

func MakeRPCCall(username string) (result string) {
	client, err := rpc.Dial("tcp", "auth-service:5001")
	if err != nil {
		log.Println("while creating RPC client: ", err)
		return
	}

	rpcPayload := RPCPayload{
		User: username,
	}

	// the response from the RPC call

	err = client.Call("RPCServer.GetLoginURL", rpcPayload, &result)
	if err != nil {
		log.Println("while calling RPC: ", err)
		return
	}

	log.Println("response from RPC: ", result)

	return result

}
