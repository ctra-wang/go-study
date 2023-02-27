build:
	go build  -o go-chat

dev:
	go run main.go

clear :
	rm -rf go-chat

add-%:
	cobra-cli add $*

build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o chatGpt.exe main.go


.PHONY: build dev clear