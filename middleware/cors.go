package middleware

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func CORS(next http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedHeaders([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)(next)
}
