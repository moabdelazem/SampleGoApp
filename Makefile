run:
	@air -c .air.toml
build:
	@go build -o main main.go	

test:
	@go test -v ./...
