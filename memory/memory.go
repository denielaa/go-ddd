package memory

import (
	"fmt"
	"sync"

	"github.com/denielaa/go-ddd/aggregate"
	"github.com/denielaa/go-ddd/domain/customer"
	"github.com/google/uuid"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mr *MemoryRepository) Get(uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}

	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Add(aggregate.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()

		if _, ok := mr.customers[id]; ok {
			return fmt.Errorf("customer already exists: %w", customer.ErrFailToCreateCustomer)
		}

		mr.Lock()
		mr.customers[c.GetID()] = c
		mr.Unlock()
		return nil
}

func (mr *MemoryRepository) Update(aggregate.Customer) error {
	if _, ok := mr.customers[c.GetID()]; ok {
		mr.Lock()
		mr.customers[c.GetID()] = c
		mr.Unlock()
		return nil
	}

	return fmt.Errorf("customer not found: %w", customer.ErrUpdateCustomer)
}