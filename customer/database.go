package customer

import "github.com/google/uuid"

var customer1 = Customer{
	ID:        uuid.New(),
	Name:      "John Doe",
	Role:      "Customer",
	Email:     "johndoe@example.com",
	Phone:     5550199,
	Contacted: true,
}
var customer2 = Customer{
	ID:        uuid.New(),
	Name:      "Jane Smith",
	Role:      "Admin",
	Email:     "janesmith@admin.com",
	Phone:     5550198,
	Contacted: false,
}
var customer3 = Customer{
	ID:        uuid.New(),
	Name:      "Alice Johnson",
	Role:      "User",
	Email:     "alice@user.com",
	Phone:     5550197,
	Contacted: true,
}

var customers = []Customer{
	customer1,
	customer2,
	customer3,
}
