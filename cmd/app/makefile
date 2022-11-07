GOCMD=go
GOTEST=$(GOCMD) test
BINARY_NAME=main.out

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

all: build
 
build:
	go build -o ${BINARY_NAME} main.go

gen-api:
#    TODO: implement me
	
lint: 
	#    TODO: implement me

test:
	go test -v main.go
 
run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}
 
clean:
	go clean
	rm ${BINARY_NAME}
