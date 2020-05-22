package storefront

import (
	"github.com/dudleycodes/golang-microservice-structure/pkg/dtos"
)

// GetMake returns a Make DTO
func (bkr Broker) GetMake(id string) (dtos.Make, error) {
	return dtos.Make{
		Value: "Some Make",
	}, nil
}
