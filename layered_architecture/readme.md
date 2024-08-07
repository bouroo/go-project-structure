# Layered Architecture

## Structure
```
├── cmd
│   ├── myapp1
│   │   └── main.go
|   └── ...
├── config
│   └── config.yaml
├── internal
│   ├── config
│   │   ├── config.go
│   │   └── config_test.go
│   ├── database
│   │   ├── database.go
│   │   └── database_test.go
│   ├── routers
│   │   ├── router.go
│   │   └── router_test.go
│   ├── handlers
│   │   ├── user
│   │   │   ├── user.go
│   │   │   └── user_test.go
│   │   └── ...
│   ├── usecases
│   │   ├── user
│   │   │   ├── user.go
│   │   │   └── user_test.go
│   │   └── ...
│   └── ...
├── Dockerfile
├── docker-compose.yaml
├── .dockerignore
└── test
    ├── ...
```

## Explanation
- `internal`: Encapsulates all project logic.
  - `config`: Handles application configuration using Viper.
  - `database`: Manages database interactions using GORM.
  - `routers`: Defines the application routes and their corresponding handlers.
  - `handlers`: Contains Fiber handlers responsible for request processing and response generation.
  - `usecases`: Provides business logic implementation, interacting with domain and database.
- `cmd`: Contains the main entry point of the application.
- `config`: Stores the application configuration file.
- `test`: Houses all tests for internal components.

## Advantages
- **Clear layering**: Provides a well-defined structure for separating concerns.
- **Testability**: Each layer can be tested independently.
- **Scalability**: Allows for easier expansion and maintainability.

## Disadvantages
- **Potential for tight coupling**: Layers can become tightly coupled if not carefully designed.