APPNAME=gocd-ci
VERSION=0.1.0-dev

build:
	go build -o ${APPNAME} .

build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -X main.Version=${VERSION}" -v -o ${APPNAME}-linux-amd64 .

build-mac:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -X main.Version=${VERSION}" -v -o ${APPNAME}-darwin-amd64 .

build-all: build-mac build-linux

all: setup
	build
	install

setup:
	go get github.com/ashwanthkumar/go-gocd
	go get github.com/spf13/cobra
	go get github.com/op/go-logging
	go get gopkg.in/yaml.v2
	# Test deps
	go get github.com/stretchr/testify

test:
	go test -v -cover github.com/ind9/gocd-ci

test-only:
	go test -v -cover github.com/ind9/gocd-ci/${name}

install: build
	sudo install -d /usr/local/bin
	sudo install -c ${APPNAME} /usr/local/bin/${APPNAME}

uninstall:
	sudo rm /usr/local/bin/${APPNAME}
