PACKAGES ?= $(shell go list ./... | grep -v /vendor/)
GENERATE ?= github.com/harborapp/harbor-go/harbor

all: clean deps test

clean:
	go clean -i ./...

deps:
	go get -u github.com/golang/lint/golint
	go get -u github.com/govend/govend
	govend -v

vendor:
	govend -vtlu

generate:
	go get -u github.com/vektra/mockery/...
	go generate $(GENERATE)

fmt:
	go fmt $(PACKAGES)

vet:
	go vet $(PACKAGES)

lint:
	for PKG in $(PACKAGES); do golint $$PKG || exit 1; done;

test:
	for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

.PHONY: all clean deps vendor generate fmt vet lint test
