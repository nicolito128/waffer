test:
	go test v ./...

run:
	go run main.go

build:
	go build -o bin/ main.exe

start:
	go test -v ./...
	go run main.go