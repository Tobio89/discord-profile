package main

import (
	"log"

	rpccontracts "discord-profile/lib/rpc-contracts"
)

type RPCServer struct{}

type RPCPayload = rpccontracts.Payload
type RPCLoginPayload = rpccontracts.LoginPayload
type RPCLoginResponse = rpccontracts.LoginResponse
type RPCSignupPayload = rpccontracts.SignupPayload
type RPCSignupResponse = rpccontracts.SignupResponse
type RPCTokenCheckPayload = rpccontracts.TokenCheckPayload
type RPCTokenCheckResponse = rpccontracts.TokenCheckResponse
type JWTValidationPayload = rpccontracts.JWTValidationPayload
type JWTValidationResponse = rpccontracts.JWTValidationResponse

func (r *RPCServer) GetLoginURL(payload RPCPayload, resp *string) error {

	log.Println("auth: received request from: ", payload.User)

	*resp = "http://localhost:5173/" + payload.User

	return nil
}

func (r *RPCServer) RequestLogin(payload RPCLoginPayload, resp *RPCLoginResponse) error {
	return app.HandleLoginRequest(payload, resp)
}

func (r *RPCServer) RequestSignup(payload RPCSignupPayload, resp *RPCSignupResponse) error {
	return app.HandleSignupRequest(payload, resp)
}

func (r *RPCServer) CheckToken(payload RPCTokenCheckPayload, resp *RPCTokenCheckResponse) error {
	return app.HandleTokenCheckRequest(payload, resp)
}

func (r *RPCServer) ValidateJWT(payload JWTValidationPayload, resp *JWTValidationResponse) error {
	return app.HandleJWTValidationRequest(payload, resp)
}
