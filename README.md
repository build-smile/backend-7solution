# Backend 7Solution

Backend 7Solution is a RESTful API for managing users, built with Go and MongoDB.

## Features

- User registration, login, update, and deletion
- JWT authentication
- RESTful endpoints
- MongoDB integration

## Getting Started

### Prerequisites

- Go 1.20+
- Docker (for MongoDB)
- Git

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/build-smile/backend-7solution.git
    cd backend-7solution
    ```

2. Start MongoDB with Docker Compose:
    ```bash
    docker compose up -d
    ```

3. Install Go dependencies:
    ```bash
    go mod tidy
    ```

4. Run the application:
    ```bash
    go run main.go
    ```

## API Endpoints

- `POST /register` — Register a new user
- `POST /login` — Login and receive JWT
- `GET /user/:id` — Get user by ID
- `PATCH /user/:id` — Update user
- `DELETE /user/:id` — Delete user
- `GET /users` — List all users

## Running Tests

To run unit tests:
```bash
go test ./...
```

## import  Postman
 seven.postman_collection.json

