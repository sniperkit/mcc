NAME := mcc
VERSION := v0.9.2
CONFIG_SCHEMA_VERSION := v1.0.0
SRCS      := $(shell find . -name '*.go' -type f)
LDFLAGS   := -ldflags "-X main.Version=$(VERSION) -X main.ConfigSchemaVersion=$(CONFIG_SCHEMA_VERSION)"

run:
	go run $(LDFLAGS) *.go -c _example/example1.yml

erd:
	go-erd -path ./dashboard/ |dot -Tsvg > ./tmp/dashboard_erd.svg

fmt:
	gofmt -s -w ./

clean:
	rm -rf _build/ release/

build:
	glide install
	mkdir -p _build
	gox $(LDFLAGS) -osarch="linux/amd64 darwin/amd64 linux/386 darwin/386" -output="_build/{{.OS}}_{{.Arch}}_{{.Dir}}"

release:
	mkdir release
	go get github.com/progrium/gh-release/...
	cp _build/* release
	gh-release create qmu/$(NAME) $(VERSION)

.PHONY: build
