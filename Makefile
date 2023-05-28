run:
	@go run ./...

test:
	@go test ./...

race:
	@go test ./... --race

build:
	@go build -o package