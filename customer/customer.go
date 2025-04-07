package customer

type Customer struct {
	ID        uint16
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
	c.ID = customerIDCounter
	customers = append(customers, c)
	customerIDCounter++
}

func DeleteCustomer(id uint16) {
	for i, customer := range customers {
		if customer.ID == id {
			customers = append(customers[:i], customers[i+1:]...)
			return
		}
	}
}

func GetCustomerByID(id uint16) *Customer {
	for _, customer := range customers {
		if customer.ID == id {
			return &customer
		}
	}
	return nil
}

func UpdateCustomer(id uint16, updatedCustomer Customer) bool {
	for i, customer := range customers {
		if customer.ID == id {
			customers[i] = updatedCustomer
			return true
		}
	}
	return false
}
