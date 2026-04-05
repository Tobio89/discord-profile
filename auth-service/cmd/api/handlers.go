package main

import (
	"crypto/rand"
	"discord-profile/auth-service/data"
	"discord-profile/auth-service/magiclink"
	b64 "encoding/base64"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func handleLoginRequest(payload RPCLoginPayload, resp *string) error {
	log.Println("auth: received login request for user: ", payload.Username)

	secureHash := make([]byte, 32)
	rand.Read(secureHash)

	sEnc := b64.StdEncoding.EncodeToString([]byte(secureHash))

	*resp = "http://localhost:5173/" + "?user=" + payload.Username + "&id=" + payload.ID + "&token=" + string(sEnc)

	return nil

}

func (app *Config) HandleLoginRequest(payload RPCLoginPayload, resp *RPCLoginResponse) error {
	log.Println("auth: received login request for user: ", payload.Username)

	// Check if user already exists in the database
	dbUser, err := app.Repo.GetUserByDiscordID(payload.ID)

	// Error performing database lookup
	if err != nil {
		log.Println("Error checking for existing user: ", err)
		*resp = RPCLoginResponse{
			Error:   true,
			Success: false,
			Message: "error checking for existing user",
			URL:     "",
		}
		return err
	}

	// User is not in the database, so we can't log them in
	if dbUser == nil {
		*resp = RPCLoginResponse{
			Error:   false,
			Success: false,
			Message: "user does not exist",
			URL:     "",
		}
		return nil
	}

	// Prepare token for magic link
	token, err := magiclink.IssueToken(app.TokenPepper)
	if err != nil {
		log.Println("Error issuing token: ", err)
		*resp = RPCLoginResponse{
			Error:   true,
			Success: false,
			Message: "error issuing token",
			URL:     "",
		}
		return err
	}

	expiry := time.Now().Add(15 * time.Minute)

	// Insert magic link into database, replacing any existing magic links for this user
	_, err = app.Repo.ReplaceMagicLinkForUser(data.MagicLink{
		DiscordUserID: payload.ID,
		TokenHash:     token.TokenHash,
		ExpiresAt:     expiry,
	})

	if err != nil {
		log.Println("Error inserting magic link: ", err)
		*resp = RPCLoginResponse{
			Error:   true,
			Success: false,
			Message: "error inserting magic link",
			URL:     "",
		}
		return err
	}

	*resp = RPCLoginResponse{
		Error:   false,
		Success: true,
		Message: "login successful",
		URL:     "http://localhost:5173/login" + "?token=" + token.RawToken,
	}

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

func (app *Config) HandleTokenCheckRequest(payload RPCTokenCheckPayload, resp *RPCTokenCheckResponse) error {

	decodedTokenBytes, err := magiclink.DecodeToken(payload.Token)
	if err != nil {
		log.Println("Error decoding token: ", err)
		resp.UserID = ""
		resp.JWT = ""
		resp.Message = "invalid token"
		return nil
	}

	hashedToken := magiclink.HashTokenBytes(decodedTokenBytes, app.TokenPepper)
	log.Println("auth: received token check request for token: ", payload.Token)

	userID, err := app.Repo.ConsumeMagicLink(hashedToken)

	if err != nil {
		log.Println("Error consuming magic link: ", err)
		resp.UserID = ""
		resp.JWT = ""
		resp.Message = "error consuming magic link"
		return err
	}

	if userID == "" {
		log.Println("Invalid or expired token: ", payload.Token)
		resp.UserID = ""
		resp.JWT = ""
		resp.Message = "invalid or expired token"
		return nil
	}

	jwt, err := app.issueJWTForUser(userID)
	if err != nil {
		log.Println("Error issuing JWT: ", err)
		resp.UserID = ""
		resp.JWT = ""
		resp.Message = "error issuing JWT"
		return err
	}

	log.Println("Successfully consumed magic link for user ID: ", userID)
	resp.UserID = userID
	resp.JWT = jwt
	resp.Message = "token valid"

	return nil
}

func (app *Config) issueJWTForUser(userID string) (string, error) {

	key := []byte(app.JWTKey)
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
		"iss": "auth-service",
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return t.SignedString(key)

}
