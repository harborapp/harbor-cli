IMPORT := github.com/umschlag/umschlag-go

PACKAGES ?= $(shell go list ./... | grep -v /vendor/)
GENERATE ?= github.com/umschlag/umschlag-go/umschlag

all: clean deps test

update:
	@which govend > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/govend/govend; \
	fi
	govend -vtlu --prune

clean:
	go clean -i ./...

generate:
	@which mockery > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/vektra/mockery/...; \
	fi
	go generate $(GENERATE)

fmt:
	go fmt $(PACKAGES)

vet:
	go vet $(PACKAGES)

lint:
	@which golint > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/golang/lint/golint; \
	fi
	for PKG in $(PACKAGES); do golint -set_exit_status $$PKG || exit 1; done;

test:
	for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

build:
	go build ./umschlag

.PHONY: all update clean generate fmt vet lint test build
