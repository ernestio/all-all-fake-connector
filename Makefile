install:
	go install -v

build:
	go build -v ./...

deps: dev-deps
	go get -u github.com/nats-io/nats
	go get -u github.com/ernestio/ernest-config-client

dev-deps:
	go get -u github.com/golang/lint/golint


lint:
	golint ./...
