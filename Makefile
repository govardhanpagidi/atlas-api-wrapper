.PHONY: build test clean
tags=logging callback metrics scheduler
cgo=0
#goos=linux
#goarch=amd64

goos=darwin
goarch=amd64

DEBUG_FLAGS=all=-N -l
CFNREP_GIT_SHA?=$(shell git rev-parse HEAD)
LINKER_FLAGS=-s -w -X github.com/govardhanpagidi/atlas-api-helper/util.defaultLogLevel=debug -X github.com/govardhanpagidi/atlas-api-helper/version.Version=${CFNREP_GIT_SHA}

build:
	@echo "==> Building handler binary"
	env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -ldflags="$(LINKER_FLAGS)" -tags="$(tags)" -o atlas_api_helper ./main.go

debug:
	@echo "==> Building handler binary for debugging"
	cfn generate
	env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -gcflags="$(DEBUG_FLAGS)" -ldflags="$(LINKER_FLAGS)" -tags="$(tags)" -o bin/handler cmd/main.go


clean:
	rm -rf bin

