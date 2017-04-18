IMPORT := github.com/umschlag/umschlag-go

PACKAGES ?= $(shell go list ./... | grep -v /vendor/)
SOURCES ?= $(shell find . -name "*.go" -type f -not -path "./vendor/*")
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
	gofmt -s -w $(SOURCES)

.PHONY: vet
vet:
	go vet $(PACKAGES)

.PHONY: misspell
misspell:
	@which misspell > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/client9/misspell/cmd/misspell; \
	fi
	misspell $(SOURCES)

.PHONY: generate
generate:
	@which mockery > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/vektra/mockery/...; \
	fi
	go generate $(GENERATE)

.PHONY: staticcheck
staticcheck:
	@which staticcheck > /dev/null; if [ $$? -ne 0 ]; then \
		go get honnef.co/go/staticcheck/cmd/staticcheck; \
	fi
	staticcheck $(PACKAGES)

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

.PHONY: unconvert
unconvert:
	@which unconvert > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/mdempsky/unconvert; \
	fi
	unconvert $(PACKAGES)

.PHONY: interfacer
interfacer:
	@which interfacer > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/mvdan/interfacer/cmd/interfacer; \
	fi
	interfacer $(PACKAGES)

.PHONY: ineffassign
ineffassign:
	@which ineffassign > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/gordonklaus/ineffassign; \
	fi
	ineffassign .

.PHONY: dupl
dupl:
	@which dupl > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/mibk/dupl; \
	fi
	dupl .

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
