package middleware

import (
	"net/http"

	"github.com/dudleycodes/golang-microservice-structure/app/storefront-api/webserver"

	"github.com/rs/zerolog/log"
)

// Authentication middleware
func Authentication(srv webserver.Server) func(h http.Handler) http.Handler {
	if srv == nil {
		log.Fatal().Msg("a nil dependency was passed to auth middleware")
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authentication")

			if !srv.ValidateJWT(token) {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("401 Unauthorized"))

				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
