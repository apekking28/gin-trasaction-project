# Gin Transaction Project

This is a simple project Transaction using the Gin framework.

## Requirements

- Go 1.15+
- Git

## Installation

1. Clone this repository:
   ```sh
   git clone https://github.com/apekking28/gin-trasaction-project.git
   cd gin-trasaction-project
   ```

2. Initialize Go modules and tidy dependencies:
   ```sh
   go mod tidy
   ```
3. Initialize Go modules and tidy dependencies:
    ```sh
    go mod tidy
    ```

4. Make sure configuration for database connection at `account_manager/database/database.go` & `payment_manager/database/database.go`

## Running the Project

To run the project, use the following command:

```sh
go run account_manager/main.go
go run payment_manager/main.go
```

The service `account_manager` will run on `http://localhost:8080`.

The service `payment_manager` will run on `http://localhost:8081`.

## Project Structure

```
.
├── account_manager
│   ├── main.go
│   ├── controllers
│   │   └── account_controller.go
│   ├── models
│   │   ├── account.go
│   │   ├── transaction.go
│   │   └── user.go
│   ├── middlewares
│   │   └── auth.go
│   └── database
│       └── database.go
├── payment_manager
│   ├── main.go
│   ├── controllers
│   │   └── transaction_controller.go
│   ├── models
│   │   └── transaction.go
│   └── database
│       └── database.go
├── go.mod
└── go.sum

```

## API Documentation

1. Import file `postman-documentation.json` to Postman
2. After that, good luck
