build:
	mkdir -p functions
	export GO111MODULE=on
	go get ./...
	go build -o functions/hook ./line-bot/server.go
