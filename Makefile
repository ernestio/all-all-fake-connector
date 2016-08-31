install:
	go install -v

build:
	go build -v ./...

deps: dev-deps
	go get github.com/nats-io/nats
	go get github.com/ernestio/ernest-config-client

dev-deps:
	go get github.com/golang/lint/golint


lint:
	golint ./...
