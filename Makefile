OUT := web-store

VERSION := $(shell git describe --always --long --dirty)

GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)

all: run

server:
	go build -i -v -o ${OUT} -ldflags="-X main.version=${VERSION}"

lint:
	@for file in ${GO_FILES} ;  do \
		golint $$file ; \
	done

clean:
	-@rm ${OUT} ${OUT}-v*

.PHONY: run server static vet lint