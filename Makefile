.PHONY: build test run docker-build docker-run


BINARY_NAME=server
DOCKER_IMAGE_NAME=task-service

build:
	go build -o ./bin/$(BINARY_NAME) ./cmd/server

test:
	go test -v ./...

run:
	go run ./cmd/server/main.go

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

docker-run:
	docker run -p 8080:8080 $(DOCKER_IMAGE_NAME)

clean:
	go clean
	rm -f ./bin/$(BINARY_NAME)

deps:
	go mod download

lint:
	golangci-lint run


all: deps build test