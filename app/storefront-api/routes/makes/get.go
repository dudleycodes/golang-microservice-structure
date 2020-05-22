package makes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dudleycodes/golang-microservice-structure/app/storefront-api/webserver"
	"github.com/gorilla/mux"

	"github.com/rs/zerolog/log"
)

// Get a make
func Get(srv webserver.Server) http.HandlerFunc {
	if srv == nil {
		log.Fatal().Msg("a nil dependency was passed to the `/makes/{makeID}` route")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		makeID := strings.TrimSpace(params["makeID"])

		makeDTO, err := srv.GetMake(makeID)
		if err != nil {

			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 Not Found"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("got make: %s", makeDTO.Value)))
	}
}
