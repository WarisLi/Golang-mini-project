# Go Mini Project: User & Product Services

## Overview
This is a mini project implemented in Go (Golang) with two services: `user` and `product`. The project follows **Hexagonal Architecture** for better separation of concerns, scalability, and testability. We use `Fiber` handling HTTP requests and defining routes for user and product services., `GORM` for interacting with the PostgreSQL database, and `Swagger` for API documentation.

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
3. API documentation using Swagger.

---

## Project Structure
This project is organized following **Hexagonal Architecture**, which separates the core domain logic from external systems like frameworks, databases, or APIs.

primary adapter -> primary port -> business logic -> secondary port -> secondary adapter


```
.
├── adapters --> adapter
│   ├── gorm_adapter.go  -> secondary adapter
│   └── http_adapter.go  -> primary adapter
├── core --> business logic, port
│   ├── service_name.go -> schema (struct)
│   ├── service_name_repository.go -> secondary port (interface)
│   └── service_name_service.go -> primary port (interface), business logic (function)
├── main.go

```

