CLEANENV := github.com/ilyakaznacheev/cleanenv
POSTGRESQL := github.com/jackc/pgx github.com/jackc/pgx/v5/pgxpool
JWT := github.com/golang-jwt/jwt/v5
VALIDATOR := github.com/go-playground/validator/v10
GIN := github.com/gin-gonic/gin
GOOSE := github.com/pressly/goose/v3/cmd/goose@latest

REST_API_BIN := build/rest_api
REST_API_SRC := cmd/rest_api/rest_api.go

GOOSE_DRIVER := postgres
GOOSE_DBHOST := 
GOOSE_DBPORT := 
GOOSE_DBNAME := 
GOOSE_DBUSER := 
GOOSE_DBPASSWORD := 
GOOSE_DBSTRING := postgresql://$(GOOSE_DBNAME):$(GOOSE_DBPASSWORD)@$(GOOSE_DBHOST):$(GOOSE_DBPORT)/$(GOOSE_DBNAME)?sslmode=disable

all: build run 

build: clean
	go build -o $(REST_API_BIN) $(REST_API_SRC)

run:
	./$(REST_API_BIN)

goose_install:
	go install $(GOOSE)

migrations_up:
	goose -dir migrations $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up

migrations_down:
	goose -dir migrations $(GOOSE_DRIVER) $(GOOSE_DBSTRING) down

docker-compose-up-silent: docker-compose-stop
	sudo docker compose -f docker-compose.yml up -d

docker-compose-stop:
	sudo docker compose -f docker-compose.yml stop

docker-compose-up: docker-compose-down
	sudo docker compose -f docker-compose.yml up

docker-compose-down:
	sudo docker compose -f docker-compose.yml down

mod:
	go mod init rest_api

get:
	go get $(CLEANENV) $(POSTGRESQL) $(JWT) $(VALIDATOR) $(GIN) $(VALIDATOR)

clean:
	rm -rf $(REST_API_BIN)