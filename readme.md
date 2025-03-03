# Go Mini Project: User & Product Services

## Overview
This is a mini project implemented in Go (Golang) with two services: **user** and **product**. The project follows **Hexagonal Architecture** for better separation of concerns, scalability, and testability. I use `Fiber` handling HTTP requests and defining routes for user and product services, `GORM` for interacting with the PostgreSQL database, `Kafka` for event streming between services, and `Swagger` for API documentation.

- **Notification Service** (https://github.com/WarisLi/Golang-notification-service.git) :
  Handles event-based notifications triggered by the Product Service.  

- **Shared Event** (https://github.com/WarisLi/Golang-shared-event.git) :
  Shared repository containing event definitions used for internal service communication.  

---
## Technologies Used

   - **Go** : Programming language
   - **Hexagonal architecture** : Architecture pattern:
   - **Fiber** : HTTP framework
   - **GORM** : ORM for database operations
   - **PostgreSQL** : Relational database
   - **Kafka** : Event streaming
   - **Swagger** : API documentation

---

## Features
1. **User Service**
   - User registration
   - Login
2. **Product Service**
   - Get product
   - Create new product
   - Update product
   - Delete product

---

## Project Structure
```
.
/project-root
│── /cmd                 # Entry point of the application
│   ├── main.go
│── /internal            # Internal code that should not be imported externally
│   ├── /core            # Business logic
│   │   ├── /ports       # Interfaces (Ports) such as Repository, Service
│   │   │   ├── product_repository.go
│   │   │   ├── product_service.go
│   │   │   ├── user_repository.go
│   │   │   ├── user_service.go
│   │   ├── /models      # Structs for entities
│   │   │   ├── product.go
│   │   │   ├── user.go
│   ├── /adapters        # Infrastructure (Database, API, HTTP)
│   │   ├── /database    # Database Adapter (GORM, SQL)
│   │   │   ├── gorm_adapter.go
│   │   ├── /http        # HTTP Adapter (Fiber)
│   │   │   ├── router.go           # Setup routes for Fiber
│   │   │   ├── product_handler.go  # HTTP handler for Product
│   │   │   ├── user_handler.go     # HTTP handler for User
│   │   │   ├── /middleware
│   │   │   │   ├── jwt_middleware.go     # JWT Middleware
│   │   │   │   ├── logging_middleware.go # Logging Middleware
│   │   ├── /producer       # Producer Adapter (Kafka)
│   │   │   ├── kafka_producer.go
│   ├── /config
│   │   ├── postgres.go  # Setup DB Connection
│   ├── /tests           # Unit tests
│   │   ├── product_test.go
│   │   ├── user_test.go
│   │   ├── utils.go
│── go.mod

```


