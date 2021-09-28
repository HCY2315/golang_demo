PROJECT:=go-ssh
GO_LDFLAGS=-ldflags " -w"

default: help
.PHONY: build

build-lux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${GO_LDFLAGS} -o ${PROJECT} main.go
build-mac: 
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ${GO_LDFLAGS} -o ${PROJECT} main.go
build-win: 
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ${GO_LDFLAGS} -o ${PROJECT}.exe main.go

help:
	@echo ""
	@echo "Build Usage:"
	@echo "    ..............................................." 
	@echo "\033[35m    make [build-mac / build-win / build-lux]  \033[0m"
	@echo "    ..............................................." 
	@echo "\033[32m    build-mac          \033[0m" "\033[36m to build binary program under Mac OS platform. \033[0m"
	@echo "\033[32m    build-lux          \033[0m" "\033[36m to build binary program under Linux platform.   \033[0m"
	@echo "\033[32m    build-win          \033[0m" "\033[36m to build binary program under Windows platform.   \033[0m"
	@echo ""