## Clean architecture
```
app/
├─ main.go              # Application entry point

assets/
├─ logs/                   # Logging assets or setup

configs/
├─ configs.go           # Configuration management

modules/
├─ servers/                # Server-related components
│  ├─ server.go         # Server initialization and setup
│  ├─ handler.go        # HTTP or other protocol handlers

├─ entities/               # Core domain entities
│  ├─ users.go          # User entity definition
│  ├─ response.go       # Common response types

├─ users/                  # User-related functionality
│  ├─ controllers/         # HTTP or CLI controllers
│  │  ├─ users_controllers.go
│  ├─ usecases/            # Application use cases
│  │  ├─ users_usecases.go
│  ├─ repositories/        # Data access interfaces and implementations
│  │  ├─ users_repositories.go

pkg/
├─ databases/              # Database-related components
│  ├─ migrations/          # Database migrations
│  ├─ postgresql.go      # PostgreSQL database setup

├─ middlewares/            # Middleware components

├─ utils/                  # Utility functions and helpers
│  ├─ connection_url_builder.go

tests/
├─ users/                  # Tests related to user functionality
│  ├─ users_test.go
 config.yaml            # Configuration file for the application

```
### Hexagonal architecture 
```
myproject/
│
├── config.yaml            # Configuration file for the application
├── main.go                # Application entry point
├── go.mod                  # Go module file
├── go.sum                  # Go module checksum file
│
├── handler/               # Adapters for handling user input
│   ├── account.go         # Account-related HTTP handlers
│   ├── customer.go        # Customer-related HTTP handlers
│   └── handler.go         # Common handler functions or setup
│
├── errs/                  # Error handling and custom errors
│   └── errs.go            # Definitions for custom errors
│
├── logs/                  # Logging setup and utilities
│   └── logs.go            # Logging configuration and functions
│
├── mock/                  # Mock implementations for testing
│   ├── mock_repository/   # Mocks for repository interfaces
│   │   ├── mock_account_repository.go
│   │   └── mock_customer_repository.go
│   └── mock_service/      # Mocks for service interfaces
│       ├── mock_account_service.go
│       └── mock_customer_service.go
│
├── repository/            # Implementations for data access
│   ├── account.go         # Repository interfaces and implementations for 
│   ├── account_db.go      # Database implementation for account
│   ├── customer.go        # Repository interfaces and implementations for 
│   ├── customer_db.go     # Database implementation for customer
│   └── customer_mock.go   # Mock implementation for customer repository
│
└── service/               # Business logic and application services
    ├── account.go         # Business logic related to account
    ├── account_service.go # Implementation of account services
    ├── customer.go        # Business logic related to customer
    ├── customer_service.go# Implementation of customer services
    └── customer_test.go   # Unit tests for customer services
```

By utilizing a core library,ensure that these functions are standardized and reusable, promoting consistency and reducing duplication across microservices.which consist of logging,kafka initialize produce,consume,read config function from local yaml file,constants,error handling, web client (resty)

```
core/
├─ config/                 # Configuration management
│  ├─ config.go          # Configuration loading and management
│
├─ constants/              # Centralized constant values
│  ├─ constants.go       # Definitions for constants used across the project
│
├─ http/                   # HTTP client and related functionality
│  ├─ client.go          # Unified client interface
│  ├─ resty_client.go    # Resty-specific client implementation
│  └─ client_test.go     # Unit tests for HTTP client
│
├─ kafka/                  # Kafka producer and consumer
│  ├─ producer.go        # Kafka producer functionality
│  ├─ consumer.go        # Kafka consumer functionality
│  └─ kafka_test.go      # Unit tests for Kafka components
│
├─ logging/                # Logging utilities
│  ├─ logger.go          # Logging setup and utilities
│  └─ logger_test.go     # Unit tests for logging
│
├─ error/                  # Custom error types and handling
│  ├─ error.go           # Custom error types and handling
│  └─ error_test.go      # Unit tests for error handling
│
├─ utils/                  # Common utility functions
│  ├─ utils.go           # Utility functions
│  └─ utils_test.go
```