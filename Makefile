.PHONY: all mod-download fmt fmt-diff ci-lint lint vet test install-tools

all: fmt-diff ci-lint lint vet test

mod-download:
	go mod download

fmt:
	goimports -w .

fmt-diff:
	test -z $$(goimports -l .) || (goimports -d . && exit 1)

ci-lint:
	golangci-lint run

lint:
	golint -set_exit_status ./...

vet:
	go vet ./...

test:
	go test -count=1 -race ./...

install-tools:
	GO111MODULE=off go get -u golang.org/x/lint/golint golang.org/x/tools/cmd/goimports github.com/golangci/golangci-lint/cmd/golangci-lint
