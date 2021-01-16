GOOS=linux
GOARCH=amd64

deps:
	go mod download

build:
	go build -o bin/deluge-automanage cmd/deluge-automanage/main.go

install: build
	cp bin/deluge-automanage /usr/local/bin