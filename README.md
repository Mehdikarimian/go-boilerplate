# go-boilerplate



# Boilerplate
### File Structure
    ..
    ├── docs                                            # Document for swagger.
    ├── src                                             # 
    │   ├── common                                      # Common Types And Struct.
        │        └── controller.go                      # Base Controller Structure Type.
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