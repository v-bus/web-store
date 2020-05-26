OUT := web-store
PKG := github.com/v_bus/rebrainme/dev/dev-06/go/web-store
VERSION := $(shell git describe --always --long --dirty)

PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)

all: run

server:
	go build -i -v -o ${OUT} -ldflags="-X main.version=${VERSION}"

vet:
	@go vet ${PKG_LIST}

lint:
	@for file in ${GO_FILES} ;  do \
		golint $$file ; \
	done

static: vet lint
	go build -i -v -o ${OUT}-v${VERSION} -tags netgo -ldflags="-extldflags \"-static\" -w -s -X main.version=${VERSION}" ${PKG}

clean:
	-@rm ${OUT} ${OUT}-v*

.PHONY: run server static vet lint
