export GO111MODULE=on
GO_RUN_VERSION := $(shell go version | awk '{print $$3}' | tr -d 'go')

ifeq ($(OS), )
	OS := $(shell uname -s | tr 'A-Z' 'a-z')
endif

all: binary

goversioncheck:
	@if ! go help mod >/dev/null 2>&1 ; then echo "Current installed Golang (=$(GO_RUN_VERSION)) doesn't have gomod support!" ; exit 1 ; fi

clean:
	rm -rf bin/

vans-%: goversioncheck
	CGO_ENABLED=0 GOOS=$* GOARCH=amd64 go build -o bin/vans -v main.go

binary: vans-$(OS)

.PHONY: all clean binary goversioncheck
