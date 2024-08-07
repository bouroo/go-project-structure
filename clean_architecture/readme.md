# Clean Architecture

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
│   ├── domain
│   │   ├── user
│   │   │   ├── user.go
│   │   │   └── user_test.go
│   │   └── ...
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
  - `domain`: Defines core business logic and entities.
  - `handlers`: Contains Fiber handlers responsible for request processing and response generation.
  - `usecases`: Provides business logic implementation, interacting with domain and database.
- `cmd`: Contains the main entry point of the application.
- `config`: Stores the application configuration file.
- `test`: Houses all tests for internal components.

## Advantages
- **Testability**: Each layer is isolated and can be tested independently.
- **Scalability**: Well-defined separation of concerns allows for easier expansion and maintainability.
- **Reusability**: Domain logic is reusable across different applications.

## Disadvantages
- **Complexity**: Increased code organization can be complex for smaller projects.