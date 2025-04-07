package customer

var customer1 = Customer{
	ID:        1,
	Name:      "John Doe",
	Role:      "Customer",
	Email:     "johndoe@example.com",
	Phone:     5550199,
	Contacted: true,
}
var customer2 = Customer{
	ID:        2,
	Name:      "Jane Smith",
	Role:      "Admin",
	Email:     "janesmith@admin.com",
	Phone:     5550198,
	Contacted: false,
}
var customer3 = Customer{
	ID:        3,
	Name:      "Alice Johnson",
	Role:      "User",
	Email:     "alice@user.com",
	Phone:     5550197,
	Contacted: true,
}

// counter for customer IDs
var customerIDCounter uint16 = 4

var customers = []Customer{
	customer1,
	customer2,
	customer3,
}
