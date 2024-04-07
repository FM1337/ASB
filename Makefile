.DEFAULT_GOAL := build

export PROJECT := "ASB"
export PACKAGE := "github.com/FM1337/ASB"


VERSION=$(shell git describe --tags --always --abbrev=0 --match=v* 2> /dev/null | sed -r "s:^v::g" || echo 0)
VERSION_FULL=$(shell git describe --tags --always --dirty --match=v* 2> /dev/null | sed -r "s:^v::g" || echo 0)

$(info $(shell mkdir -p $(DIRS)))
BIN=$(CURDIR)/bin
export GOBIN=$(CURDIR)/bin

build: go-fetch go-build
	@echo

go-prepare: go-fetch
	go generate -x ./...

go-fetch:
	go mod download
	go mod tidy

go-upgrade-deps:
	go get -u ./...
	go mod tidy

go-upgrade-deps-patch:
	go get -u=patch ./...
	go mod tidy

go-dlv: go-prepare
	dlv debug \
		--headless --listen=:2345 \
		--api-version=2 --log \
		--allow-non-terminal-interactive \
		${PACKAGE} -- --debug

go-debug: go-prepare
	go run ${PACKAGE} --debug

go-debug-fast:
	go run ${PACKAGE} --debug

go-build: go-prepare go-fetch
	CGO_ENABLED=0 \
	go build \
		-ldflags '-d -s -w -extldflags=-static' \
		-tags=netgo,osusergo,static_build \
		-installsuffix netgo \
		-buildvcs=false \
		-trimpath \
		-o ${PROJECT} \
		${PACKAGE}
