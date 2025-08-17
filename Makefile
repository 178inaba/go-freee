.PHONY: all fmt fmt-diff lint lint-fix vet test go-install-tools

all: fmt lint-fix lint vet test

fmt:
	goimports -w .

fmt-diff:
	test -z $$(goimports -l .) || (goimports -d . && exit 1)

lint:
	docker compose run --rm lint

lint-fix:
	docker compose run --rm lint --fix

vet:
	go vet ./...

test:
	go test -race -cover ./...

go-install-tools:
	go install golang.org/x/tools/cmd/goimports@latest
