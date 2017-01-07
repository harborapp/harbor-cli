IMPORT := github.com/umschlag/umschlag-go

PACKAGES ?= $(shell go list ./... | grep -v /vendor/)
GENERATE ?= $(IMPORT)/umschlag

.PHONY: all
all: clean deps test

.PHONY: update
update:
	@which govend > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/govend/govend; \
	fi
	govend -vtlu --prune

.PHONY: clean
clean:
	go clean -i ./...

.PHONY: fmt
fmt:
	go fmt $(PACKAGES)

.PHONY: vet
vet:
	go vet $(PACKAGES)

.PHONY: generate
generate:
	@which mockery > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/vektra/mockery/...; \
	fi
	go generate $(GENERATE)

.PHONY: errcheck
errcheck:
	@which errcheck > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/kisielk/errcheck; \
	fi
	errcheck $(PACKAGES)

.PHONY: varcheck
varcheck:
	@which varcheck > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/opennota/check/cmd/varcheck; \
	fi
	varcheck $(PACKAGES)

.PHONY: structcheck
structcheck:
	@which structcheck > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/opennota/check/cmd/structcheck; \
	fi
	structcheck $(PACKAGES)

.PHONY: lint
lint:
	@which golint > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/golang/lint/golint; \
	fi
	for PKG in $(PACKAGES); do golint -set_exit_status $$PKG || exit 1; done;

.PHONY: test
test:
	for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

.PHONY: check
check: test

.PHONY: build
build:
	go build -v ./umschlag
