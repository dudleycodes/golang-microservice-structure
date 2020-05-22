package storefront

import (
	"github.com/dudleycodes/golang-microservice-structure/pkg/dtos"
)

// GetModel returns a Model DTO
func (bkr Broker) GetModel(id string) (dtos.Model, error) {
	return dtos.Model{
			Value: "some model",
		},
		nil
}
