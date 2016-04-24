.PHONY: all install vet test run

all: run

install: vet test
	@go install ./cmd/...

vet:
	@go vet ./...

test:
	@go test ./...

run: install
	@templatetest