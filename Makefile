.PHONY: build
build:
	go build -o bin/slack-timer cmd/slack-timer/main.go

.PHONY: format
format:
	go mod tidy
	goimports -w .

.PHONY: test
test:
	go test -v ./...
