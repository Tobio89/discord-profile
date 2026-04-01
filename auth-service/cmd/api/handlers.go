package main

import (
	"crypto/rand"
	"discord-profile/auth-service/data"
	"discord-profile/auth-service/magiclink"
	b64 "encoding/base64"
	"log"
)

func handleLoginRequest(payload RPCLoginPayload, resp *string) error {
	log.Println("auth: received login request for user: ", payload.Username)

	secureHash := make([]byte, 32)
	rand.Read(secureHash)

	sEnc := b64.StdEncoding.EncodeToString([]byte(secureHash))

	*resp = "http://localhost:5173/" + "?user=" + payload.Username + "&id=" + payload.ID + "&token=" + string(sEnc)

	return nil

}

func (app *Config) HandleLoginRequest(payload RPCLoginPayload, resp *string) error {
	log.Println("auth: received login request for user: ", payload.Username)

	// Check if user already exists in the database
	dbUser, err := app.Repo.GetUserByDiscordID(payload.ID)

	if err != nil {
		log.Println("Error checking for existing user: ", err)
		return err
	}

	// If user already exists, return response indicating that
	if dbUser == nil {
		*resp = ""
		return nil
	}

	token, err := magiclink.IssueToken(app.TokenPepper)

	*resp = "http://localhost:5173/login" + "?token=" + token.RawToken

	return nil

}

func (app *Config) HandleSignupRequest(payload RPCSignupPayload, resp *RPCSignupResponse) error {
	log.Println("auth: received signup request for user: ", payload.Username)

	// Check if user already exists in the database
	dbUser, err := app.Repo.GetUserByDiscordID(payload.ID)

	if err != nil {
		log.Println("Error checking for existing user: ", err)
		return err
	}

	// If user already exists, return response indicating that
	if dbUser != nil {
		log.Println("User already exists with discord ID: ", payload.ID)
		resp.AlreadyExists = true
		resp.Message = "User already exists"
		return nil
	}

	// If user does not exist, create new user in the database
	resp.AlreadyExists = false
	resp.Message = "User created successfully"

	userID, err := app.Repo.InsertDiscordUser(data.DiscordUser{
		DiscordUserID: payload.ID,
	})

	if err != nil {
		log.Println("Error inserting new user: ", err)
		return err
	}

	log.Println("Successfully inserted new user with ID: ", userID)
	return nil

}
