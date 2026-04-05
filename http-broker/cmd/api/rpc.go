package main

import (
	rpccontracts "discord-profile/lib/rpc-contracts"
	"log"
	"net/rpc"
)

type RPCTokenCheckPayload = rpccontracts.TokenCheckPayload
type RPCTokenCheckResponse = rpccontracts.TokenCheckResponse

func DialRPCServer() (*rpc.Client, error) {
	client, err := rpc.Dial("tcp", "auth-service:5001")
	if err != nil {
		return nil, err
	}

	return client, nil

}

func RPCRequestTokenValidation(rawToken string) (result string, err error) {
	client, err := DialRPCServer()
	if err != nil {
		log.Println("while dialing RPC server: ", err)
		return "", err
	}

	payload := RPCTokenCheckPayload{
		Token: rawToken,
	}

	response := RPCTokenCheckResponse{}

	err = client.Call("RPCServer.RequestSignup", payload, &response)
	if err != nil {
		log.Println("while calling RPC: ", err)
		return "", err
	}

	log.Println("response from RPC: ", result)

	return result, nil
}
