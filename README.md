# CRM Backend

## Tech Stack
* Go-Lang

## Clone the project

```bash
git clone https://github.com/swapab/crm-backend.git
cd crm-backend
```

## Build
```bash
go build
```

## Run test
```bash
go run test
```

## Start Server
```bash
go run main.go
```

## API Endpoints

### 1. Create a new customer
- **URL**: `/customers`
- **Method**: `POST`
- **Request Body**: 
  - `name`: string (required)
  - `email`: string (required)
  - `phone`: string (required)
  - `role`: string (required)
  - `Connected`: boolean (optional)

- **Response**:
  - `201 Created`: Customer created successfully
  - `400 Bad Request`: Invalid input data
  - `500 Internal Server Error`: Server error

### 2. Get all customers
- **URL**: `/customers`
- **Method**: `GET`
- **Response**:
  - `200 OK`: List of customers
  - `500 Internal Server Error`: Server error

### 3. Get a customer by ID
- **URL**: `/customers/{id}`
- **Method**: `GET`
- **Response**:
  - `200 OK`: Customer details
  - `404 Not Found`: Customer not found
  - `500 Internal Server Error`: Server error

### 4. Update a customer by ID
- **URL**: `/customers/{id}`
- **Method**: `PUT`
- **Request Body**: 
  - `name`: string (optional)
  - `email`: string (optional)
  - `phone`: string (optional)
  - `role`: string (optional)
  - `Connected`: boolean (optional)
- **Response**:
- `200 OK`: Customer updated successfully
  - `400 Bad Request`: Invalid input data
  - `404 Not Found`: Customer not found
  - `500 Internal Server Error`: Server error

### 5. Delete a customer by ID
- **URL**: `/customers/{id}`
- **Method**: `DELETE`
- **Response**:
  - `204 No Content`: Customer deleted successfully
  - `404 Not Found`: Customer not found
  - `500 Internal Server Error`: Server error
