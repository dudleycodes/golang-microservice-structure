package webserver

import (
	"github.com/dudleycodes/golang-microservice-structure/internal/storefront"
	"github.com/dudleycodes/golang-microservice-structure/pkg/authentication"
)

// Config for the storefront API.
type Config struct {
	Auth       authentication.Config
	Storefront storefront.Config

	Port int
}

func validateConfig(cfg Config) error {
	// This function would be used to validate a hydrated configuration; return an error if its invalid.
	return nil
}
