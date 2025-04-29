# Go FastHTTP Example

A high-performance RESTful API server example built with Go and FastHTTP. This project demonstrates how to build a scalable web service using the FastHTTP framework, implementing clean architecture principles with proper separation of concerns.

## Features

- High-performance HTTP server using FastHTTP
- Clean architecture with proper separation of concerns
- RESTful API endpoints for user management
- In-memory repository implementation
- Comprehensive Makefile for common development tasks

## Prerequisites

- Go 1.21 or higher
- Make (optional, for using Makefile commands)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/captain-corgi/go-fasthttp-example.git
   cd go-fasthttp-example
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

## Usage

The project includes a Makefile with common development commands:

```bash
# Build the application
make build

# Run the application
make run

# Run tests
make test

# Run linter
make lint

# Clean build artifacts
make clean
```

Alternatively, you can run the server directly:

```bash
go run cmd/server/main.go
```

By default, the server starts on port 8080. You can specify a different port using the `-port` flag:

```bash
go run cmd/server/main.go -port 3000
```

## API Endpoints

### User Management

- **GET /users**
  - Retrieves all users
  - Response: 200 OK (with array of user data)

- **GET /users/{id}**
  - Retrieves a user by ID
  - Response: 200 OK (with user data) or 404 Not Found

- **POST /users**
  - Creates a new user
  - Request body: JSON user object
  - Response: 201 Created (with created user data)

- **PUT /users/{id}**
  - Updates an existing user
  - Request body: JSON user object
  - Response: 200 OK (with updated user data)

- **DELETE /users/{id}**
  - Deletes a user
  - Response: 204 No Content

## Project Structure

```
.
├── cmd/
│   └── server/          # Application entry point
├── internal/
│   ├── domain/          # Domain models and interfaces
│   ├── handler/         # HTTP request handlers
│   ├── repository/      # Data access layer
│   └── service/         # Business logic layer
├── Makefile            # Build and development commands
└── README.md           # Project documentation
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.