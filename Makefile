VERSION?="0.1.0"

PROJECTNAME=$(shell basename "$(PWD)")

.PHONY: all
all:
	lint
	build
	test

.PHONY: build
build:
	echo "Building $(PROJECTNAME)"
	GOOS=windows GOARCH=386 go build -o bin/main.exe cmd/goblog/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-arm64 cmd/goblog/main.go

.PHONY: run
run:
	echo "Starting $(PROJECTNAME)"
	go run cmd/goblog/main.go

.PHONY: docs
docs:
	cd cmd; cd goblog; swag init .../cmd/goblog

.PHONY: test
test:
	echo "Testing $(PROJECTNAME)"	
	go test ./cmd/goblog/... -v
	# ginkgo ./cmd/goblog/

.PHONY: lint
lint:
	golangci-lint run

