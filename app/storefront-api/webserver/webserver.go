package webserver

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/dudleycodes/golang-microservice-structure/internal/storefront"
	"github.com/dudleycodes/golang-microservice-structure/pkg/authentication"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

// Server exposes all functionalities of the Storefront API.
type Server interface {
	authentication.Authentication
	storefront.Storefront
}

// Broker manages the internal state of the Storefront API.
type Broker struct {
	authentication.Authentication
	storefront.Storefront

	cfg    Config      // the api service's configuration
	router *mux.Router // the api service's route collection
}

// New initializes a new Storefront API.
func New(cfg Config) (*Broker, error) {
	r := &Broker{}

	err := validateConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}
	r.cfg = cfg

	r.Authentication, err = authentication.New(cfg.Auth)
	if err != nil {
		return nil, fmt.Errorf("invalid auth configuration: %w", err)
	}

	r.Storefront, err = storefront.New(cfg.Storefront)
	if err != nil {
		return nil, fmt.Errorf("invalid storefront configuration: %w", err)
	}

	// Do other setup work here...

	return r, nil
}

// Start the Storefront service
func (bkr *Broker) Start(binder func(s Server, r *mux.Router)) {
	bkr.router = mux.NewRouter().StrictSlash(true)
	binder(bkr, bkr.router)

	// Do other startup work here...
	l, err := net.Listen("tcp", ":"+strconv.Itoa(bkr.cfg.Port))
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to bind to TCP port %d for listening.", bkr.cfg.Port)
		os.Exit(13)
	} else {
		log.Info().Msgf("Starting webserver on TCP port %04d", bkr.cfg.Port)
	}

	if err := http.Serve(l, bkr.router); errors.Is(err, http.ErrServerClosed) {
		log.Warn().Err(err).Msg("Web server has shut down")
	} else {
		log.Fatal().Err(err).Msg("Web server has shut down unexpectedly")
	}
}

// PingDependencies ping all Storefront API dependencies.
func (bkr *Broker) PingDependencies(failFast bool) error {
	if !bkr.Storefront.Ping() {
		return fmt.Errorf("Couldn't ping storefront dependencies")
	}

	return nil
}
