.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build

.PHONY: run
run:
    gin -d ./cmd/apiserver -i --all -b "go-rest-api"
