# GoSpire: Effortless REST API Building with Go-Gin-GORM
This is a dynamic and efficient starter kit for constructing REST APIs in Go, featuring the splendid Gin framework, the GORM ORM powerhouse, and the robust PostgreSQL database. Experience API development like never before with a master-slave database architecture.

## Features

-   **Boilerplate Structure:** A well-organized project structure that promotes scalability and maintainability.
-   **RESTful Endpoints:** Ready-to-use controllers and middleware for creating user and post APIs.
-   **GRPC Support:** Controllers and middleware for handling GRPC requests. (Upcoming)
-   **Database Abstraction:** Utilizes GORM for seamless interaction with the PostgreSQL database.
-   **Master-Slave Database Support:** Implements a robust master-slave database architecture for enhanced data management and performance optimization.
-   **User Authentication:** Middleware for authentication and CORS handling.
-   **Seed Data:** Pre-configured commands for populating the database with initial data.
-   **Logging and Error Handling:** Basic logging setup with storage for logs.

## Project Structure

```
├── app
│   ├── console
│   │   ├── commands
│   │   │   ├── seed.go
│   ├── grpc
│   │   ├── controllers
│   │   │   ├── UserController.go
│   │   │   └── ... other controllers
│   │   └── middleware
│   │       ├── Authenticate.go
│   │       └── ... other middlewares
│   ├── helpers
│   │   ├── helpers.go
│   │   └── ... other helpers
│   ├── http
│   │   ├── controllers
│   │   │   ├── UserController.go
│   │   │   └── PostController.go
│   │   ├── middleware
│   │   │   ├── Authenticate.go
│   │   │   └── Cors.go
│   │   └── requests
│   │       ├── CreateUserRequest.go
│   │       ├── UpdateUserRequest.go
│   │       └── ... other requests
│   ├── interfaces
│   │   ├── UserRepositoryInterface.go
│   │   └── ... other interfaces
│   ├── repositories
│   │   ├── UserRepository.go
│   │   └── ... other repositories
│   ├── services
│   │   ├── UserService.go
│   │   └── ... other services
│   ├── models
│   │   ├── Base.go
│   │   ├── Post.go
│   │   └── User.go
│   └── providers
│       ├── RouteServiceProvider.go
│       └── ... other providers
├── bootstrap
│   ├── app
│   │   ├── app.go
│   │   └── database.go
│   ├── server
│   │   └── server.go
├── config
│   ├── app.go
│   ├── database.go
│   └── init.go
├── database
│   ├── migrations.go
│   ├── seeders
│   │   ├── UsersTableSeeder.go
│   │   └── ...
│   └── ...
├── public
│   └── ...
├── resources
│   ├── views
│   │   ├── index.html
│   │   └── ... other view files
│   └── ...
├── routes
│   ├── api.go
│   ├── grpc.go
│   └── web.go
├── storage
│   ├── logs
│   └── temp
├── tests
│   └── ...
├── .air.toml
├── .env
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── README.md
```
## Getting Started

1.  Clone this repository.
2.  Set up your environment variables in the `.env` file.
3.  Install dependencies using `go mod tidy`.
4.  Start the application with `go run main.go` or using a tool like [Air](https://github.com/cosmtrek/air) for hot-reloading.

## Usage

-   Define your routes in the `routes` directory.
-   Create models, repositories, and services as needed.
-   Add controllers and middleware in the `app/http/controllers` and `app/http/middleware` directories.
-   Customize and extend the functionality to suit your specific project requirements.


## Contributing

If you find any issues or have suggestions, feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](https://chat.openai.com/c/LICENSE).