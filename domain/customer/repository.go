package customer

import (
	"errors"

	"github.com/denielaa/go-ddd/aggregate"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound = errors.New("customer not found")
	ErrFailToCreateCustomer = errors.New("failed to create customer")
	ErrUpdateCustomer = errors.New("failed to update customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}