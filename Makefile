build:
	go build -race -o ./bin/todofinder main.go

test:
	go test ./... -race