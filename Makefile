PROJECTNAME := "puppetdb-proxy"
VERSION := $(shell cat VERSION)
BUILD := $(shell git rev-parse --short HEAD)
DESCRIPTION := "The proxy service for Puppet for store data from old Puppet agents to new PuppetDB."
MAINTAINER := "Michael Bruskov <mixanemca@yandex.ru>"

# Go source files, ignore vendor directory
GOFILES = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X=main.version=$(VERSION) -X=main.build=$(BUILD)"

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

.PHONY: build help

all: build

## build: Compile the binary.
build: vet lint clean
	@go build $(LDFLAGS) -o $(PROJECTNAME) $(GOFILES)

## linux: Compile the binary for GNU/Linux amd64.
linux: vet lint clean
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(PROJECTNAME) $(GOFILES)

## vet: Runs the Go vet command on the packages named by the import paths.
vet:
	@go vet $(GOFILES)

## lint: Runs the golint command.
lint:
	@go get -u golang.org/x/lint/golint
	@golint $(GOFILES)

## fpm-deb: Build Debian package
fpm-deb:
	fpm -s dir -t deb -n $(PROJECTNAME) -v $(VERSION) \
		--deb-priority optional --category admin \
		--force \
		--url https://github.com/nemca/$(PROJECTNAME) \
		--description $(DESCRIPTION) \
		-m $(MAINTAINER) \
		--license "Apache 2.0" \
		-a amd64  \
		$(PROJECTNAME)=/usr/bin/$(PROJECTNAME) \
		etc/=/


## clean: Cleanup.
clean:
	@-rm -f $(PROJECTNAME)
	@-rm -f *.deb

## help: Show this message.
help: Makefile
	@echo "Available targets:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
