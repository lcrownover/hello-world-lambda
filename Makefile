all: build zip

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go

zip:
	zip main.zip main
