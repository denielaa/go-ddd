package aggregate

import (
	"errors"

	"github.com/denielaa/go-ddd/entity"
	"github.com/denielaa/go-ddd/valueobject"
	"github.com/google/uuid"
)

type Customer struct {
	person				*entity.Person
	products 			[]*entity.Item
	transactions 	[]*valueobject.Transaction
}

var (
	ErrInvalidPerson = errors.New("a customer has to have a valid person")
)

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID: uuid.New(),
	}

	return Customer{
		person: person,
		products: make([]*entity.Item, 0),
		transactions: make([]*valueobject.Transaction, 0),
	}, nil
}