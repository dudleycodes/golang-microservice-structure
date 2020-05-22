package mock

import "github.com/dudleycodes/golang-microservice-structure/pkg/dtos"

// Result sets the result of a mock storefront function
type Result func(c *mockConfig)

type mockConfig struct {
	pingShouldFail bool
	getMake        error
	getModel       error
}

// Mock for mocking Storefront service
type Mock struct {
	cfg mockConfig
}

// New creates a new, mock storefront service
func New(opts ...Result) Mock {
	r := Mock{}

	for _, o := range opts {
		if o != nil {
			o(&r.cfg)
		}
	}

	return r
}

// Ping mocks the storefront's Ping() function
func (m Mock) Ping() bool {
	return !m.cfg.pingShouldFail
}

// GetMake mocks a Storefront GetMake() call
func (m Mock) GetMake(string) (dtos.Make, error) {
	if m.cfg.getMake != nil {
		return dtos.Make{}, m.cfg.getMake
	}

	return dtos.Make{
		Value: "mock make",
	}, nil
}

// GetMakeResult sets the result of the mock GetMake()
func GetMakeResult(e error) Result {
	return func(c *mockConfig) {
		c.getMake = e
	}
}

// GetModel mocks a Storefront GetModel() call
func (m Mock) GetModel(string) (dtos.Model, error) {
	if m.cfg.getModel != nil {
		return dtos.Model{}, m.cfg.getModel
	}

	return dtos.Model{
		Value: "mock model",
	}, nil
}

// GetModelResult sets the result of the mock GetMake()
func GetModelResult(e error) Result {
	return func(c *mockConfig) {
		c.getModel = e
	}
}
