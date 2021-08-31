# meta
NAME := vc-mop

export GO111MODULE=on

## Run tests
.PHONY: test
test:
	go test ./...

## Lint
.PHONY: lint
lint:
	go vet ./...
	golint -set_exit_status ./...