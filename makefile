LOGRUS := github.com/sirupsen/logrus github.com/sirupsen/logrus@v1.9.3
ROUTER := github.com/julienschmidt/httprouter
CLEANENV := github.com/ilyakaznacheev/cleanenv
POSTGRESQL := github.com/jackc/pgx github.com/jackc/pgx/v5/pgxpool
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
	go get $(POSTGRESQL)

push: 
	@read -p "inter commit string: " commit; \
	git add .; \
	git commit -m "$$commit"; \
	git push origin main;

clean:
	rm -rf app
	rm -rf app.sock
