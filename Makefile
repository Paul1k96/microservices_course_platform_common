LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

go/lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.yaml