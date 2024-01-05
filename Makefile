.PHONY: build aecli

export GO111MODULE=on

.pre-build:
	$(eval GOGC = off)
	$(eval CGO_ENABLED = 0)

.pre-aecli:
	$(eval APP_NAME = aecli)

build: aecli

aecli: .pre-build .pre-aecli
	go build -o build/aecli main.go
	sudo ln -sf $(PWD)/build/aecli /usr/local/bin

deps:
	go mod download
