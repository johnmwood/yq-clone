.PHONY:
build:
	go build -o bin/ccyq -v

.PHONY:
run:
	go run cmd/yq/main.go

.PHONY:
test:
	go test -v ./...
