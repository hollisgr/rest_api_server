# REST API SERVER

Server implement basic CRUD operations on USER entity:

type User struct { \
	Id           string `json:"id"` \
	Login        string `json:"login"` \
	FirstName    string `json:"first_name"` \
	SecondName   string `json:"second_name"` \
	Email        string `json:"email"` \
	PasswordHash string `json:"password_hash"` \
}

Based on:
- Golang (gin);
- Postgresql.

## Configure server:
### Configure **.env** file in root:

BIND_IP=`{server host}`; \
LISTEN_PORT=`{server port}`; \
PSQL_HOST=`{postgresql host}`; \
PSQL_PORT=`{postgresql port}`;  \
PSQL_DBNAME=`{postgresql db name}`;  \
PSQL_USER=`{postgresql username}`;  \
PSQL_PASSWORD=`{postgresql password}`; \
JWT_SECRET_KEY=`{your secret key}`; \
JWT_TOKEN_EXP_TIME=`{token expiration time in hours}`.

## Configure GOOSE migrations:

- Install **GOOSE**: `go install github.com/pressly/goose/v3/cmd/goose@latest`;

- Set values in **makefile**: \
    GOOSE_DBHOST := `{postgresql host}`; \
    GOOSE_DBPORT := `{postgresql port}`; \
    GOOSE_DBNAME := `{postgresql db name}`; \
    GOOSE_DBUSER := `{postgresql username}`; \
    GOOSE_DBPASSWORD := `{postgresql password}`.

- After setting values use `make migrations_up` to create table;
- Or `make migrations_down` to delete table.

## Install and run server:
- After configure you can use makefile for **quick build and run server** by using command `make all`;
- For build-only use command `make build`
- You can build it without makefile using command: \
    `go build -o rest_api cmd/rest_api/rest_api.go` \
    and run it `./rest_api`;
- Or you can use `go run cmd/rest_api/rest_api.go`.

## Docker compose:
- by default there is 2 docker container (server + database): \
server container uses 8181:8080 port by default; \
postgres container uses 25432:5432 port by default.
- `make docker-compose-up` to run constainers;
- `make docker-compose-up-silent` to run containers in detached in the background;
- `make docker-compose-stop` to stop containers;
- `make docker-compose-down` to stop containers and removes containers, networks, volumes, and images created by up.

## Routes:
### Without token:
- POST("/users") - create new user
- POST("/auth") - authorization by login and password

### Protected:
- GET("/users/:id") - get user by id
- GET("/users") - get list of all users
- PATCH("/users/:id") - update user by id
- DELETE("/users/:id") - delete user by id