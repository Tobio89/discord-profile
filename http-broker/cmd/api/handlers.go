package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type jsonValidationResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) PostValidateToken(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	if token == "" {
		app.writeJSON(w, http.StatusBadRequest, jsonValidationResponse{
			Error:   true,
			Message: "missing token",
		})
		return
	}

	rpcResponse, err := RPCRequestTokenValidation(token)
	if err != nil {
		app.writeJSON(w, http.StatusInternalServerError, jsonValidationResponse{
			Error:   true,
			Message: "error validating token",
		})
		return
	}

	if rpcResponse.UserID == "" {
		app.writeJSON(w, http.StatusBadRequest, jsonValidationResponse{
			Error:   true,
			Message: "token is invalid",
			Data:    rpcResponse,
		})
		return
	}

	data := struct {
		UserID string `json:"user_id"`
	}{
		UserID: rpcResponse.UserID,
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    rpcResponse.JWT,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // dev only; true in HTTPS/prod
		SameSite: http.SameSiteLaxMode,
		MaxAge:   60 * 60 * 24,
	})

	app.writeJSON(w, http.StatusOK, jsonValidationResponse{
		Error:   false,
		Message: "validation successful",
		Data:    data,
	})

}

func (app *Config) GetValidateToken(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	if token == "" {
		app.writeJSON(w, http.StatusOK, jsonValidationResponse{
			Error:   false,
			Message: "Get token - your token was missing",
			Data:    nil,
		})
		return

	}
	app.writeJSON(w, http.StatusOK, jsonValidationResponse{
		Error:   false,
		Message: "Get token - your token was: " + token,
		Data:    nil,
	})
}

func (app *Config) Root(w http.ResponseWriter, r *http.Request) {
	app.writeJSON(w, http.StatusOK, jsonValidationResponse{
		Error:   false,
		Message: "Welcome to the profile broker API",
		Data:    nil,
	})
}
