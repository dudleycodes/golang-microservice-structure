package storefront

import (
	"fmt"

	"github.com/dudleycodes/golang-microservice-structure/pkg/dtos"
)

// Storefront exposes all functionalities of the Storefront service.
type Storefront interface {
	GetMake(string) (dtos.Make, error)
	GetModel(string) (dtos.Model, error)
	Ping() bool
}

// Broker manages the internal state of the Storefront service.
type Broker struct {
	cfg Config // the storefront's configuration
}

// New initializes a new Storefront service.
func New(cfg Config) (*Broker, error) {
	r := &Broker{}

	if err := validateConfig(cfg); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return r, nil
}

// Ping checks to see if the storefront's database is responding.
func (brk *Broker) Ping() bool {
	// This function would check the storefront's dependencies (datastores and whatnot); useful for Kubernetes probes
	return true
}
