build:
	mkdir -p functions
	go get ./...
	go build -o functions/hook ./line-bot/server.go
