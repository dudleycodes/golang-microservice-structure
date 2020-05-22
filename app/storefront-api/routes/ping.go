package routes

import (
	"net/http"

	"github.com/dudleycodes/golang-microservice-structure/app/storefront-api/webserver"
	"github.com/rs/zerolog/log"
)

// Ping is for the Kubernetes liveness probe
func Ping(srv webserver.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if _, err := w.Write([]byte(`"PONG"`)); err != nil {
			log.Fatal().Msg("Something went very, very wrong")
		}
	}
}
