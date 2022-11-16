GOCMD=go
GOTEST=$(GOCMD) test
BINARY_NAME=main.out
DB_URL=postgres://root:pass@localhost:5432/digdb?sslmode=disable

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

all: build
 
build:
	cd cmd/app \
	&& go build -o ${BINARY_NAME} main.go

gen-api:
	cd specs \
	&& oapi-codegen \
		--config ./config.yml \
		--package=specs \
		-o ./contract.gen.go \
		contract.openapi.yaml
	
lint: 
	golangci-lint run

test:
	go test -v main.go
 
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

migrate-seeds-up:
	migrate -source file://migrations/postgresql/seeds -database "${DB_URL}" up

migrate-seeds-down:
	migrate -source file://migrations/postgresql/seeds -database "${DB_URL}" down
