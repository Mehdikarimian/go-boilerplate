# go-boilerplate

A great starting point for building RESTful APIs in Go using Gin framework, connecting to a PostgreSQL database with both pure and gorm, MongoDB database, Redis database.

### Features

-   Implements the Clean Architecture pattern for a scalable and maintainable
-   Uses the Gin framework for efficient and fast handling of HTTP requests
-   Integrates with Gorm TypeOrm for powerful and flexible database operations (https://gorm.io/)
-   Integrates with MongoDb database
-   Integrates with Redis database
-   Uses Go Swag for Generating Rest Docs and Swagger panel (https://github.com/swaggo/swag)  
-   Uses Air for live reload app (https://github.com/cosmtrek/air)

##### Authentication

-   Supports JWT authentication with configurable expiration and issuer, allowing for flexible and secure authentication processes.

### Getting Started

##### Prerequisites

-   Go version 1.18.1 or higher

To get up and running with the Go-Boilerplate, follow these simple steps:

```
$ git clone https://github.com/Mehdikarimian/go-boilerplate
$ cd go-boilerplate
$ cp internal/config/.env.example internal/config/.env # create a copy of the example environment file, and also follow configuration steps on the difference section below
$ go src/main.go Or air
```

#### Configuration

# Boilerplate
### File Structure
    ..
    ├── docs                                            # Document for swagger.
    ├── src                                             # 
    │   ├── common                                      # Common Types And Struct.
    |   │        └── controller.go                      # Base Controller Structure Type.
    │   │        └── model.go                           # Base Model Structure Type.
    |   ├── config                                      # Configs
    |   |        └── config.go                          # Base Config Module and Env Init.
    │   └── controllers                                 # Controllers
    │   │         └── articles.controllers.go           # Article Controller (example).
    │   │         └── base.go                           # Base Controller Structure.
    │   │         └── products.controllers.go           # Products Controller (example).
    │   │         └── swagger.controllers.go            # Swagger Controller
    │   │         └── users.controllers.go              # Users Article Controller (example).
    │   │                                               #
    │   ├── core                                        # Core Configures
    │   │   └── db                                      # Db Configures 
    │   │       └── gorm.go                             # Gorm (Golang Typeorm) Configure File 
    │   │       └── mongo.go                            # MongoDb Configure File
    │   │       └── postgres.go                         # PostgresSQL Configure File
    │   │       └── redis.go                            # Redis Configure File
    │   │                                               #
    │   ├── models                                      # Models
    │   │   └── base.go                                 # Base Model Structure.
    │   │   └── article.model.go                        # Article Model (example).
    │   │   └── cache.model.go                          # Cache Model (example).
    │   │   └── product.model.go                        # Product Model (example).
    │   │   └── user.model.go                           # User Model (example).
    │   │                                               #
    │   ├── utils                                       # Utils.
    │   │   └── http.go                                 # Http Utils
    │   │                                               #
    │   ├── main.go                                     # Main File.
    ├── .env.example                                    # 
    ├── Dockerfile                                      #
    ├── docker-compose.yml                              #
    ├── .air.toml                                       #
    └── ...