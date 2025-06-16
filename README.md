# REST API SERVER

### Server implement basic CRUD operations on USER entity:
```golang
type User struct { 
	Id           string `json:"id"` 
	Login        string `json:"login"` 
	FirstName    string `json:"first_name"` 
	SecondName   string `json:"second_name"` 
	Email        string `json:"email"` 
	PasswordHash string `json:"password_hash"` 
}
```
### Based on:
- **Golang (gin)**;
- **Postgresql**.

## Configure server:
- Configure **.env** file in root:
```
BIND_IP={server host}
LISTEN_PORT={server port}
PSQL_HOST={postgresql host}
PSQL_PORT={postgresql port}
PSQL_DBNAME={postgresql db name}
PSQL_USER={postgresql username}
PSQL_PASSWORD={postgresql password}
JWT_SECRET_KEY={your secret key}
JWT_TOKEN_EXP_TIME={token expiration time in hours}.
```

## Configure GOOSE migrations:

- Install **GOOSE**: `go install github.com/pressly/goose/v3/cmd/goose@latest`;

- Set values in **makefile**:
```
GOOSE_DBHOST := {postgresql host}
GOOSE_DBPORT := {postgresql port}
GOOSE_DBNAME := {postgresql db name}
GOOSE_DBUSER := {postgresql username}
GOOSE_DBPASSWORD := {postgresql password}.
```
- After setting values use `make migrations_up` to create table;
- Or `make migrations_down` to delete table.

## Build and run server:
- After configure you can use makefile for **quick build and run server** by using command `make all`
- For **build-only** use command `make build`
- You can **build it without makefile** using command: 
- - `go build -o rest_api cmd/rest_api/rest_api.go` 
- - and run it `./rest_api`
- Or you can use `go run cmd/rest_api/rest_api.go`

## Docker compose:
- by default there is 2 docker container (server + database): 
- - server container uses **8181:8080** port by default; 
- - postgres container uses **25432:5432** port by default.
- `make docker-compose-up` to run constainers;
- `make docker-compose-up-silent` to run containers in detached in the background;
- `make docker-compose-stop` to stop containers;
- `make docker-compose-down` to stop containers and removes containers, networks, volumes, and images created by up.

## Routes:
### POST("/users") - create new user
```
REQUEST:
  method: POST
  headers: 
    Content-Type: "application/json"
  body: 
    {
        "login":"{your_login}",
        "password":"{your_password}",
        "first_name":"{your_first_name}",
        "second_name":"{your_second_name}",
        "email":"{your@email.com}"
	}
RESPONSE:
  status: 200
  body: 
	{
  	  "message": "created user with id: {your_id}",
	  "status": "OK",
	  "success": true 
	}

  status: 400
  body: 
	{
	  "message": "{error text}",
	  "status": "Bad Request",
	  "success": false
	}
```
### POST("/auth") - authorization by login and password
```
REQUEST:
  method: POST
  headers: 
    Content-Type: "application/json"
  body: {
	  "login":"{your_login}",
	  "password":"{your_password}",
	}
RESPONSE:
  status: 200
  body: 
	{
	  "message": "Bearer {your_token}",
	  "status": "OK",
	  "success": true
	}

  status: 400
  body: 
	{
	  "message": "{error text}",
	  "status": "Bad Request",
	  "success": false
	}
	
  status: 401
  body: 
  	{
	  "message": "{error text}",
	  "status": "Unauthorized",
	  "success": false
  	}
```
### GET("/users/:id") - get user by id
```
REQUEST:
  method: GET
  headers: 
    Authorization: "Bearer {token}"
RESPONSE:
  status: 200
  body: 
	{
	  "success": "true",
	  "user": {
		"id": {your_id},
		"login": "{your_login}",
		"first_name": "{your_first_name}",
		"second_name": "{your_second_name}",
		"email": "{your@email.com}"
  	  }
	}

status: 400
  body: 
	{
	  "message": "{error text}",
	  "status": "Bad Request",
	  "success": false
	}
	
  status: 401
  body: 
  	{
	  "message": "{error text}",
	  "status": "Unauthorized",
	  "success": false
  	}
```

### GET("/users") - get list of all users
```
REQUEST:
  method: GET
  headers: 
    Authorization: "Bearer {token}"
RESPONSE:
  status: 200
  body: 
	{
	  "success": "true",
	  "users list": 
	  [
	  	{
		  "id": {your_id},
		  "login": "{your_login}",
		  "first_name": "{your_first_name}",
		  "second_name": "{your_second_name}",
		  "email": "{your@email.com}"
	  	},
	  	{
		  "id": {your_id},
		  "login": "{your_login}",
		  "first_name": "{your_first_name}",
		  "second_name": "{your_second_name}",
		  "email": "{your@email.com}"
	  	},
	  	{
		  "id": {your_id},
		  "login": "{your_login}",
		  "first_name": "{your_first_name}",
		  "second_name": "{your_second_name}",
		  "email": "{your@email.com}"
	  	}
	  ]
	}

  status: 400
  body: 
	{
	  "message": "{error text}",
	  "status": "Bad Request",
	  "success": false
	}
	
  status: 401
  body: 
  	{
	  "message": "{error text}",
	  "status": "Unauthorized",
	  "success": false
  	}
```

### PATCH("/users/:id") - update user by id

```
REQUEST:
  method: PATCH
  headers: 
    Authorization: "Bearer {token}"
    Content-Type: "application/json"
  body: 
	{
	  "login":"{your_new_login}",
	  "first_name":"{your_new_first_name}",
	  "second_name":"{your_new_second_name}",
	  "email":"{your_new@email.com}"
	}
RESPONSE:
  status: 200
  body: 
	{
	  "success": "true",
	  "user": {
		"id": {your_id},
		"login": "{your_login}",
		"first_name": "{your_first_name}",
		"second_name": "{your_second_name}",
		"email": "{your@email.com}"
	  }
	}

  status: 400
  body: 
	{
	  "message": "{error text}",
	  "status": "Bad Request",
	  "success": false
	}
	
  status: 401
  body: 
  	{
	  "message": "{error text}",
	  "status": "Unauthorized",
	  "success": false
  	}
```

### DELETE("/users/:id") - delete user by id

```
REQUEST:
  method: DELETE
  headers:
    Authorization: "Bearer {token}"
RESPONSE:
  status: 200
  body: 
	{
	  "message": "deleted user_id: {id}",
	  "status": "OK",
	  "success": true
	}

  status: 400
  body: 
	{
	  "message": "{error text}",
	  "status": "Bad Request",
	  "success": false
	}
	
  status: 401
  body: 
  	{
	  "message": "{error text}",
	  "status": "Unauthorized",
	  "success": false
  	}
```