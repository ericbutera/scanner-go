IMAGE_NAME=storage-api
MAIN=main.go

.PHONY: build
build:
	go build -o bin/app

.PHONY: clean
clean:
	rm bin/*

.PHONY: docker-build
docker-build:
	docker build -t ${IMAGE_NAME} .

.PHONY: docker-run
docker-run:
	docker run --rm -p 8080:8080 ${IMAGE_NAME}

.PHONY: help
help:
	echo "Usage: make [target]"

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test:
	go test ./...

.PHONY: run
run:
	go run ${MAIN}

.PHONY: docs
dogs:
	swag init

.PHONY: protoc
protoc:
	protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protobuf/*.proto