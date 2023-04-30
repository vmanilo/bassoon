GOBINPATH=$(shell go env GOPATH)/bin
VERSION=0.0.1

fmt:
	gofmt -w -s .

test:
	@go test ./...

test-cover:
	@go test ./... -coverprofile cover.out ./...
	@go tool cover -func cover.out | grep total | awk '{print $3}'

html-coverage: test
	go tool cover -html=cover.out

build:
	go build -ldflags "-X main.version=${VERSION}" -o bassoon cmd/app/main.go

run: build
	./bassoon

lint: tools
	@$(GOBINPATH)/golangci-lint run -c golangci.yml ./...

tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

mocks:
	mockgen -package=mocks -source internal/app/service/service.go -destination internal/app/service/mocks/mock.go
