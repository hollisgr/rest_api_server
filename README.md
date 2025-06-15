# REST API SERVER

Based on:
- Golang (gin)
- Postgresql

## Configure server:
Configure ".env" file in root:

BIND_IP=`{host}` \
LISTEN_PORT=`{port}` \
PSQL_HOST=`{postgresql host}` \
PSQL_PORT=`{postgresql port}`  \
PSQL_DBNAME=`{postgresql db name}`  \
PSQL_USER=`{postgresql username}`  \
PSQL_PASSWORD=`{postgresql password}` \
JWT_SECRET_KEY=`{your secret key}` \
JWT_TOKEN_EXP_TIME=`{token expression time in hours}`

## Configure GOOSE migrations

- Install GOOSE: `go install github.com/pressly/goose/v3/cmd/goose@latest`

- Set values in MAKEFILE: \
    GOOSE_DBHOST := `{postgresql host}` \
    GOOSE_DBPORT := `{postgresql port}` \
    GOOSE_DBNAME := `{postgresql db name}` \
    GOOSE_DBUSER := `{postgresql username}` \
    GOOSE_DBPASSWORD := `{postgresql password}`

- After setting values use `make migrations_up` to create table
- Or `make migrations_down` to delete table

## Install and run server:
- After configure you can use makefile for quick build server by command `make all`
- You can build it manually using command: `go build -o rest_api cmd/rest_api/rest_api.go` and run it `./rest_api`
- Or you can use `go run cmd/rest_api/rest_api.go`

## Routes:
### Without token:
- POST("/users") - create new user
- POST("/auth") - authorization by login and password

### Protected:
- GET("/users/:id") - get user by id
- GET("/users") - get list of all users
- PATCH("/users/:id") - update user by id
- DELETE("/users/:id") - delete user by id