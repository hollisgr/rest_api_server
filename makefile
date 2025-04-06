LOGRUS := github.com/sirupsen/logrus
ROUTER := github.com/julienschmidt/httprouter
CLEANENV := github.com/ilyakaznacheev/cleanenv
APP := cmd/main/app.go

all: mod get run_server

build: clean
	go build $(APP)

run_server: build
	./app

mod:
	go mod init rest_api_server

get:
	go get $(ROUTER)
	go get $(LOGRUS)
	go get $(CLEANENV)

clean:
	rm -rf app
	rm -rf app.sock

users:
	curl localhost:8080/users

user:
	curl localhost:8080/users/123

create:
	curl -X POST localhost:8080/users/123

delete:
	curl -X DELETE localhost:8080/users/123