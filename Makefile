deps:
	@go get github.com/golang/dep/cmd/dep
	@dep ensure

lint:
	@go get honnef.co/go/tools/cmd/megacheck
	@go get github.com/alecthomas/gometalinter
	@gometalinter --install > /dev/null
	@gometalinter --vendor --disable-all --enable=errcheck --enable=golint --enable=megacheck --enable=vet ./...

install:
	@go install -v

build:
	@go build -v ./...
