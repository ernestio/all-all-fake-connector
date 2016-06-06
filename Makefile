install:
	go install -v

build:
	go build -v ./...

deps: dev-deps

dev-deps:
	go get -u github.com/golang/lint/golint


lint:
	golint ./...
