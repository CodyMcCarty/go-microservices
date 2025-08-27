# Go Microservices (course project)

This repository contains a REST-based microservice built in Go, backed by a Postgres database running in Docker for local development, and follows the LinkedIn learning | [Build a microservice](https://www.linkedin.com/learning/build-a-microservice-with-go) by Frank P Moley

The goal of this project was not to create production-ready software, but to learn and apply common **patterns for microservice development in Go**:
- **HTTP request/response handling** (using the [Echo](https://echo.labstack.com/) framework)
- **Data access abstraction** (using [GORM](https://gorm.io/) for Postgres)
- The instructor emphasized this course is about *patterns*, not polished Go style or optimization.

---

## 🏗 Project Structure
- **Models**: single structs shared across DB + JSON API.
- **Database**: `DatabaseClient` interface + `Client` implementation using GORM.
- **Server**: Echo HTTP server, routes grouped by entity (`/customers`, `/products`, etc).
- **Custom errors**: `NotFoundError`, `ConflictError` → mapped to `404` and `409` in handlers.  

## 🚀 Running the Service

### Prerequisites
- Go 1.25+
- docker (optional)
- HTTPie (optional)

### Prepare your environment
All instructions below assume **WSL** (Windows Subsystem for Linux). 
If you are using native Windows or another OS, adjust paths and commands as needed.  

1. Check if the Postgres container (`local-pg`) is running:   
`docker ps`
1. If it’s not running, start it with the helper script:  
$`./dat/postgres_start.sh`
1. Exec into the running container:    
$`docker exec -it local-pg /bin/bash`
1. Launch psql from inside the docker container   
$`psql -U postgres`
1. Check for the data and schema  
$`select * from wisdom.customers limit 10;`
1. if needed, Copy/paste schema file and then data file from /dat directory into $psql
  
### Run the server

`$ go run .`  
Server will start on http://localhost:8080  
check:  
http://localhost:8080/readiness  
http://localhost:8080/liveness  
http://localhost:8080/vendors  

### Endpoints
The service exposes basic CRUD operations for four entities:  
Health checks: /readiness, /liveness   
/customers → Create, Read (all/by id), Update, Delete  
/products → Create, Read (all/by id), Delete  
/vendors → Create, Read (all/by id), Delete  
/services → Create, Read (all/by id), Delete  

### Example Requests
Examples of requests for bash and chrome dev tools can be found in `customer.go`. HTTPie or others can also be used. There's **no** postman.

## Notes
This is a learning project built for educational purposes.
It demonstrates patterns and concepts but is not production-ready.
I have a ton of personal notes scattered through the project especially in customer related files.
There were issues with his Update POST logic, I got update customer working, but I want to look more into it before doing the other endpoints.  

## Future Extensions
- multiple dbs
- main app with 1 or 2 microservices.
- config file, along with other comments in *customer
- logs to include which user updated customer and when, etc.
- config with env var (kelsey Hightower or Seth Vargo) 3.BuildDockerImage 1:30
- middleware 3.middleware. CORS
- testing, 
- swaggo swagger,
- call another service, google login,

