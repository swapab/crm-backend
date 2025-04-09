package customer

import "github.com/google/uuid"

type Customer struct {
	ID        uuid.UUID
	Name      string
	Role      string
	Email     string
	Phone     uint64
	Contacted bool
}

func GetCustomers() []Customer {
	return customers
}

func AddCustomer(c Customer) {
	// Assign a new ID to the customer
	c.ID = uuid.New()
	customers = append(customers, c)
}

func DeleteCustomer(id uuid.UUID) {
	for i, customer := range customers {
		if customer.ID == id {
			customers = append(customers[:i], customers[i+1:]...)
			return
		}
	}
}

func GetCustomerByID(id uuid.UUID) *Customer {
	for _, customer := range customers {
		if customer.ID == id {
			return &customer
		}
	}
	return nil
}

func UpdateCustomer(id uuid.UUID, updatedCustomer Customer) bool {
	for i, customer := range customers {
		if customer.ID == id {
			customers[i] = updatedCustomer
			return true
		}
	}
	return false
}
