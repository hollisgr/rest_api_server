LOGRUS := github.com/sirupsen/logrus github.com/sirupsen/logrus@v1.9.3
ROUTER := github.com/julienschmidt/httprouter
CLEANENV := github.com/ilyakaznacheev/cleanenv
MONGODB := go.mongodb.org/mongo-driver/v2/mongo go.mongodb.org/mongo-driver/bson go.mongodb.org/mongo-driver/bson/primitive  go.mongodb.org/mongo-driver/mongo go.mongodb.org/mongo-driver/mongo/options
APP := cmd/main/app.go

all: mod get build

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
	go get $(MONGODB)

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