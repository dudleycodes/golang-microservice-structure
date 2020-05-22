package authentication

//
// This would normal be a package in another repo...
//

// Auth exposes all functionalities of the Auth agent
type Authentication interface {
	ValidateJWT(token string) bool
}

// Broker manages the internal state of the Auth agent.
type Broker struct{}

// New create a new authorization agent.
func New(cfg Config) (Authentication, error) {
	return &Broker{}, nil
}

// ValidateJWT validates the JWT token against the remote authorization service.
func (bkr *Broker) ValidateJWT(token string) bool {
	// Do JWT validation stuff here
	return true
}
