package mock

import (
	"github.com/dudleycodes/golang-microservice-structure/internal/storefront"
	"github.com/dudleycodes/golang-microservice-structure/pkg/authentication"

	mockstore "github.com/dudleycodes/golang-microservice-structure/internal/storefront/mock"
	mockauth "github.com/dudleycodes/golang-microservice-structure/pkg/authentication/mock"
)

// Result tells the webserver Mock how to return a specific result
type Result func(c *mockConfig)

type mockConfig struct {
}

// Mock the webserver
type Mock struct {
	authentication.Authentication
	storefront.Storefront

	cfg mockConfig
}

// New create a new Mock webserver.
func New(opts ...Result) Mock {
	r := Mock{
		Authentication: mockauth.New(),
		Storefront:     mockstore.New(),
	}

	for _, o := range opts {
		if o != nil {
			o(&r.cfg)
		}
	}

	return r
}

// WithAuthentication attaches a customized authentication mock
func (m Mock) WithAuthentication(opts ...mockauth.Result) Mock {
	m.Authentication = mockauth.New(opts...)

	return m
}

// WithStorefront attaches a customized storefront mock
func (m Mock) WithStorefront(opts ...mockstore.Result) Mock {
	m.Storefront = mockstore.New(opts...)

	return m
}
