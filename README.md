# Fiber Hexagonal Backend

A **backend-only Go project** using **Fiber**, **Hexagonal Architecture**, and **MariaDB/GORM**.
Includes JWT authentication (cookie-based) and integration with third-party APIs.

---

## Table of Contents

* [Features](#features)
* [Project Structure](#project-structure)
* [Getting Started](#getting-started)
* [Environment Variables](#environment-variables)
* [Running the Server](#running-the-server)
* [API Endpoints](#api-endpoints)
* [JWT Authentication](#jwt-authentication)
* [Third-Party API Integration](#third-party-api-integration)
* [Testing](#testing)
* [Next Steps](#next-steps)

---

## Features

* Hexagonal Architecture (Ports & Adapters)
* Fiber REST API
* MariaDB with GORM
* In-memory repository for testing
* JWT authentication via HTTP-only cookie
* External API integration (example)
* Fully testable core domain logic

---

## Project Structure

```
fiber-hex/
├─ cmd/api/                  # Entry point
│  ├─ main.go
│  └─ bootstrap.go
├─ internal/
│  ├─ core/
│  │  ├─ domain/            # Domain entities
│  │  ├─ ports/             # Interfaces
│  │  └─ usecase/           # Business logic
│  ├─ adapters/
│  │  ├─ http/              # Fiber handlers, routes, middleware
│  │  ├─ repository/        # Repository adapters (memory, MariaDB)
│  │  └─ external/          # Third-party API adapters
├─ pkg/
│  ├─ config/               # Config loader
│  └─ logger/               # Logger
├─ .env                     # Environment variables
├─ go.mod
└─ go.sum
```

---

## Getting Started

### Prerequisites

* Go >= 1.21
* MariaDB server
* Optional: Docker for local MariaDB

---

### Environment Variables

Create `.env`:

```dotenv
APP_ENV=development
PORT=3000

DB_USER=root
DB_PASS=secret
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=fiberhex

JWT_SECRET=your_jwt_secret
```

---

### Running the Server

```bash
go mod tidy
go run ./cmd/api
```

Server runs on `http://localhost:3000`

---

## API Endpoints

### Books CRUD

| Method | Endpoint        | Description             |
| ------ | --------------- | ----------------------- |
| GET    | /api/books/     | List all books          |
| GET    | /api/books/\:id | Get book by ID          |
| POST   | /api/books/     | Create a new book (JWT) |
| PUT    | /api/books/\:id | Update book (JWT)       |
| DELETE | /api/books/\:id | Delete book (JWT)       |

### Health Check

```
GET /health
```

---

## JWT Authentication

* Issue JWT via `AuthHandler` login
* Stored as **HTTP-only cookie**
* Protect routes using `JWTMiddleware`

---

## Third-Party API Integration

* Adapter pattern for external APIs
* Example: `WeatherAPI` service
* Use ports to call external API from use case

---

## Testing

Unit tests for use case logic can run without Fiber or DB:

```bash
go test ./internal/core/usecase/...
```

---

## Next Steps

* Add more entities and relations (Author → Books)
* Dockerize backend + MariaDB
* Add OpenAPI/Swagger documentation
* Add middleware (logging, CORS, rate-limiting)
* CI/CD integration

---

## License

MIT License
