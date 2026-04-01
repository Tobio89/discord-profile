package main

import (
	"log"
	"net/rpc"

	rpccontracts "discord-profile/lib/rpc-contracts"
)

type RPCServer struct{}

type RPCPayload = rpccontracts.Payload
type RPCLoginPayload = rpccontracts.LoginPayload
type RPCLoginResponse = rpccontracts.LoginResponse
type RPCSignupPayload = rpccontracts.SignupPayload
type RPCSignupResponse = rpccontracts.SignupResponse

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {

	log.Println("broker: received request from: ", payload.User)

	response := MakeRPCCall(payload.User)

	*resp = response

	return nil
}

func (r *RPCServer) RequestLogin(payload RPCLoginPayload, resp *RPCLoginResponse) error {
	log.Println("broker: received login request for user:", payload.Username)

	response, err := RPCRequestLogin(payload)
	if err != nil {
		log.Println("error while requesting login:", err)
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

func RPCRequestLogin(loginReq RPCLoginPayload) (result RPCLoginResponse, err error) {
	client, err := DialRPCServer()
	if err != nil {
		log.Println("while dialing RPC server: ", err)
		return RPCLoginResponse{}, err
	}

	payload := RPCLoginPayload{
		Username: loginReq.Username,
		ID:       loginReq.ID,
		Token:    loginReq.Token,
	}

	err = client.Call("RPCServer.RequestLogin", payload, &result)
	if err != nil {
		log.Println("while calling RPC: ", err)
		return RPCLoginResponse{
			Success: false,
			Message: "unable to login",
			URL:     "",
		}, err
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
