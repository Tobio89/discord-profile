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
		JWT    string `json:"jwt"`
	}{
		UserID: rpcResponse.UserID,
		JWT:    rpcResponse.JWT,
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

func (app *Config) GetCheckToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		app.writeJSON(w, http.StatusUnauthorized, jsonValidationResponse{
			Error:   true,
			Message: "No session cookie found",
			Data:    nil,
		})
		return
	}

	jwt := cookie.Value
	if jwt == "" {
		app.writeJSON(w, http.StatusUnauthorized, jsonValidationResponse{
			Error:   true,
			Message: "Session cookie is empty",
			Data:    nil,
		})
		return
	}

	rpcResponse, err := RPCRequestJWTValidation(jwt)
	if err != nil {
		app.writeJSON(w, http.StatusInternalServerError, jsonValidationResponse{
			Error:   true,
			Message: "Error validating token",
		})
		return
	}
	if !rpcResponse.Valid {
		app.writeJSON(w, http.StatusUnauthorized, jsonValidationResponse{
			Error:   false,
			Message: "JWT is invalid",
			Data:    rpcResponse,
		})

	}
	app.writeJSON(w, http.StatusOK, jsonValidationResponse{
		Error:   false,
		Message: "JWT is valid",
		Data:    rpcResponse,
	})
}

func (app *Config) Root(w http.ResponseWriter, r *http.Request) {
	app.writeJSON(w, http.StatusOK, jsonValidationResponse{
		Error:   false,
		Message: "Welcome to the profile broker API",
		Data:    nil,
	})
}
