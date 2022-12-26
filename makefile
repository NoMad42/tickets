GOCMD=go
GOTEST=$(GOCMD) test
BINARY_DIR=cmd/app/
BINARY_NAME=main.out
DB_URL=postgres://root:pass@localhost:5432/digdb?sslmode=disable

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

all: build
 
build:
	go build -o ${BINARY_DIR}${BINARY_NAME} ${BINARY_DIR}main.go

gen-api:
	cd specs \
	&& oapi-codegen \
		--config ./config.yml \
		--package=specs \
		-o ./contract.gen.go \
		contract.openapi.yaml
	
lint: 
	golangci-lint run

# test:
# 	go test -v main.go
 
run:
	cd cmd/app \
	&& go build -o ${BINARY_NAME} main.go \
	&& ./${BINARY_NAME}
 
clean:
	cd cmd/app \
	&& go clean \
	&& rm ${BINARY_NAME}

migrate-up:
	migrate -source file://migrations/postgresql -database "${DB_URL}" up

migrate-down:
	migrate -source file://migrations/postgresql -database "${DB_URL}" down

migrate-refresh: 
	make migrate-down
	make migrate-up

mig-r: migrate-refresh

migrate-fresh: 
	migrate -source file://migrations/postgresql -database "${DB_URL}" drop
	make migrate-up

mig-f: migrate-fresh

psql:
	docker exec -it dive-into-go-db psql "${DB_URL}"
