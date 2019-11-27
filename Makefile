all: build

.PHONY: build
build: test
	go build -o hello main.go

.PHONY: test
test:
	go test -v ./...

