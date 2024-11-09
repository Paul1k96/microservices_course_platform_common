LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
	GOBIN=$(LOCAL_BIN) go install go.uber.org/mock/mockgen@v0.5.0

go/lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.yaml

generate: generate/go

generate/go:
	@go generate ./...