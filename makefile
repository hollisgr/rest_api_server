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

push: 
	@read -p "inter commit string: " commit; \
	git add .; \
	git commit -m "$$commit"; \
	git push origin main;

clean:
	rm -rf app
	rm -rf app.sock
