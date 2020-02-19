SHELL = /bin/bash
OUTPUTDIR = ./bin

.ONESHELL:
.PHONY: build doc clean help

all: help

## build: Build binary file.
build:
	@export GOFLAGS=-mod=vendor
	@echo Building...
	CGO_ENABLED=0 GOOS=linux go build -a -o $(OUTPUTDIR)/rota
	@echo Success

## doc: Generate document.
doc:
	@apidoc -i msg

## clean: Clean build files.
clean:
	-@go clean
	-@rm $(OUTPUTDIR)/rota

## help: Show this help.
help: Makefile
	@echo Usage: make [command]
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'