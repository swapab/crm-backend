package main

// import customer package
import (
	"crm-backend/customer"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

// This function will return a list of customers
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(customer.GetCustomers())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	// Add new customer to the database (in this case, a slice)
	var newCustomer customer.Customer
	fmt.Println("BOdy", json.NewDecoder(r.Body))
	err := json.NewDecoder(r.Body).Decode(&newCustomer)
	fmt.Println("newCustomer", newCustomer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	customer.AddCustomer(newCustomer)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newCustomer)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Println("ID", id)

	customerID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	customer.DeleteCustomer(customerID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println("ID", id)

	// Get customer by id
	customerID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid cust ID", http.StatusBadRequest)
		return
	}
	cust := customer.GetCustomerByID(customerID)
	if cust == nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(cust)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
	}

	customerID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	var updatedCustomer customer.Customer
	err = json.NewDecoder(r.Body).Decode(&updatedCustomer)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	customer.UpdateCustomer(customerID, updatedCustomer)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(updatedCustomer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func showIndexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func main() {
	router := mux.NewRouter()
	// Initialize the HTTP server
	router.HandleFunc("/", showIndexPage).Methods("GET")
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")

	// Start the server on port 3000
	fmt.Println("Starting Server...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println("Failed to start server")
		return
	}
}
