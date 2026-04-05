package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type jsonValidationResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) ValidateToken(w http.ResponseWriter, r *http.Request) {
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

	app.writeJSON(w, http.StatusOK, jsonValidationResponse{
		Error:   false,
		Message: "token is valid",
		Data:    rpcResponse,
	})

}

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// who can connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/validate-token/:token", app.ValidateToken)

	return mux
}
