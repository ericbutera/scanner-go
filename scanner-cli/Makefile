IMAGE_NAME=scanner-go
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
	docker run --rm \
	-e AWS_PROFILE=test \
	-v ${HOME}/.aws/credentials:/root/.aws/credentials:ro \
	-v ${pwd}/config:/app/config \
	${IMAGE_NAME}

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