# API-SANDBOX

A concise project description detailing what this API does. This Go application provides a robust API framework with efficient routing using the Gin framework and includes Swagger documentation for easy endpoint references.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [File Structure](#file-structure)
- [Usage](#usage)
- [Documentation](#documentation)

## Prerequisites

- Go (at least v1.15)
- Gin framework
- Swagger for Go

Install Go: [Go's official installation guide](https://golang.org/doc/install)  

Install Gin:
```go
go get -u github.com/gin-gonic/gin
```

Install Swagger for Go:
```go
go get -u github.com/swaggo/swag/cmd/swag
```

## Installation
1. Clone this repository:
```bash
git clone <repository-url>
```

2. Navigate to the project directory and install all dependencies:
```go
go mod tidy
```

## File Structure

- *main.go*: The entrypoint of the application, initializes and starts the server.
- *router.go*: Takes care of routing and HTTP request handling. It leverages the Gin framework to ensure efficient routing and middlewares.
- *handler.go*: This file houses the handler function responsible for managing the GET API endpoint, ensuring that requests are processed and the appropriate response is returned.
- *swagger.go*: Automatically generates Swagger documentation for the API. This aids in understanding the API structure, endpoints, and what each endpoint is expected to do.

## Usage
1.Start the server:
```go
go run main.go
```

Navigate to your API (default is http://localhost:8080) using a browser or a tool like Postman to test your endpoints.

## Documentation
For API documentation, navigate to /swagger (e.g., http://localhost:8080/swagger-ui/index.html) to view the Swagger UI with a list of all API endpoints, their purposes, and expected responses.
