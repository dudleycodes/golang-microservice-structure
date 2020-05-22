package mock

// Result tells the Authentication Mock how to return a specific result
type Result func(c *mockConfig)

type mockConfig struct {
	validateJWTShouldFail bool
}

// Mock the Authentication agent
type Mock struct {
	cfg mockConfig
}

// New create a new Mock authorization agent.
func New(opts ...Result) *Mock {
	r := &Mock{}

	for _, o := range opts {
		if o != nil {
			o(&r.cfg)
		}
	}

	return r
}

// ValidateJWTFail sets the result for the mock ValidateJWT() function
func ValidateJWTFail() Result {
	return func(c *mockConfig) {
		c.validateJWTShouldFail = true
	}
}

// ValidateJWT mocks JWT authentication results
func (m Mock) ValidateJWT(token string) bool {
	return !m.cfg.validateJWTShouldFail
}
