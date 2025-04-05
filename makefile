all: mod get_router run_server

run_server:
	go run cmd/main/app.go

mod:
	go mod init rest_api_server

get_router:
	go get github.com/julienschmidt/httprouter

users:
	curl localhost:8080/users

user:
	curl localhost:8080/users/123

create:
	curl -X POST localhost:8080/users/123

delete:
	curl -X DELETE localhost:8080/users/123