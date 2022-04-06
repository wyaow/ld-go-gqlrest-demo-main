.PHONY: all gen build run test lint

all: gen build test lint

gen:
	@# https://github.com/golang/go/issues/44129 // add "-mod=mod" as workaround for go 1.16 bug
	go run -mod=mod github.com/speedoops/go-gqlrest

build:
	go build

release:
	go build -o go-gqlrest-federation.exe main.go
	
run:	
	# @command -v air &>/dev/null || go install github.com/cosmtrek/air
	# air
	go run main.go

test:
	#go test -timeout 30s -run ^TestTodo ./... -v -failfast -coverpkg=github.com/speedoops/go-gqlrest/... -coverprofile=coverage.out
	go test -timeout 30s -run ^TestTodo ./... -v -failfast -coverpkg=./...,github.com/speedoops/go-gqlrest/... -coverprofile=coverage.out
	go tool cover -func=coverage.out | tail -n 1 | awk '{print "Total coverage: " $$3 " of statements"}'
	go tool cover -html=coverage.out -o coverage.html

smoke:
	go test -gcflags=all=-l -timeout 30s ./... -short -v -coverprofile=coverage.out
	go tool cover -func=coverage.out | tail -n 1 | awk '{print "Total coverage: " $$3 " of statements"}'

lint:
	golangci-lint run --timeout=5m
