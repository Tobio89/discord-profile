package rpc

import (
	"log"
	"net/rpc"

	rpccontracts "discord-profile/lib/rpc-contracts"
)

type RPCPayload = rpccontracts.Payload
type RPCLoginPayload = rpccontracts.LoginPayload
type RPCLoginResponse = rpccontracts.LoginResponse

type LoginRequest struct {
	Username string
	ID       string
	Token    string
}

type RPCSignupPayload = rpccontracts.SignupPayload
type RPCSignupResponse = rpccontracts.SignupResponse

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

func RPCRequestLogin(loginReq RPCLoginPayload) (result RPCLoginResponse, err error) {
	client, err := DialRPCServer()
	if err != nil {
		log.Println("while dialing RPC server: ", err)
		return RPCLoginResponse{Success: false, Message: "unable to connect to auth"}, err
	}

	err = client.Call("RPCServer.RequestLogin", loginReq, &result)
	if err != nil {
		log.Println("while calling RPC: ", err)
		return RPCLoginResponse{Success: false, Message: "unable to login"}, err
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
