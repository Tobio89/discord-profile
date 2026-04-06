package main

import (
	rpccontracts "discord-profile/lib/rpc-contracts"
	"log"
	"net/rpc"
)

type RPCTokenCheckPayload = rpccontracts.TokenCheckPayload
type RPCTokenCheckResponse = rpccontracts.TokenCheckResponse

type RPCJWTValidationPayload = rpccontracts.JWTValidationPayload
type RPCJWTValidationResponse = rpccontracts.JWTValidationResponse

func DialRPCServer() (*rpc.Client, error) {
	client, err := rpc.Dial("tcp", "auth-service:5001")
	if err != nil {
		return nil, err
	}

	return client, nil

}

func RPCRequestTokenValidation(rawToken string) (result RPCTokenCheckResponse, err error) {
	client, err := DialRPCServer()
	if err != nil {
		log.Println("while dialing RPC server: ", err)
		return rpccontracts.TokenCheckResponse{UserID: "", Message: "error validating token"}, err
	}

	payload := RPCTokenCheckPayload{
		Token: rawToken,
	}

	err = client.Call("RPCServer.CheckToken", payload, &result)
	if err != nil {
		log.Println("while calling RPC: ", err)
		return rpccontracts.TokenCheckResponse{UserID: "",
			Message: "error validating token",
		}, err
	}

	log.Println("response from RPC: ", result)

	return result, nil
}

func RPCRequestJWTValidation(jwt string) (result RPCJWTValidationResponse, err error) {
	client, err := DialRPCServer()
	if err != nil {
		log.Println("while dialing RPC server: ", err)
		return RPCJWTValidationResponse{UserID: "", Valid: false, Message: "failed to dial RPC"}, err
	}

	payload := RPCJWTValidationPayload{
		Token: jwt,
	}

	err = client.Call("RPCServer.ValidateJWT", payload, &result)
	if err != nil {
		log.Println("while calling RPC: ", err)
		return RPCJWTValidationResponse{
			UserID:  "",
			Message: "error validating token",
			Valid:   false,
		}, err
	}

	log.Println("response from RPC: ", result)

	return result, nil
}
