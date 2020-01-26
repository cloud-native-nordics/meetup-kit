PROJECT=github.com/cloud-native-nordics/meetup-kit
GO_VERSION=1.13.1
BINARIES=meetup-kit
CACHE_DIR = $(shell pwd)/bin/cache

all: build
build: $(BINARIES)

.PHONY: $(BINARIES) 
$(BINARIES):
	make shell COMMAND="make bin/$@"

.PHONY: bin/meetup-kit
bin/meetup-kit: bin/%: vendor
	CGO_ENABLED=0 go build -mod=vendor -ldflags "$(shell ./hack/ldflags.sh)" -o bin/$* ./cmd/$*

shell:
	mkdir -p $(CACHE_DIR)/go $(CACHE_DIR)/cache
	docker run -it --rm \
		-v $(CACHE_DIR)/go:/go \
		-v $(CACHE_DIR)/cache:/.cache/go-build \
		-v $(shell pwd):/go/src/${PROJECT} \
		-w /go/src/${PROJECT} \
		-u $(shell id -u):$(shell id -g) \
		golang:$(GO_VERSION) \
		$(COMMAND)

vendor:
	go mod tidy
	go mod vendor

tidy: /go/bin/goimports
	go mod tidy
	go mod vendor
	gofmt -s -w pkg cmd
	goimports -w pkg cmd
	go run hack/cobra.go

/go/bin/goimports:
	go get golang.org/x/tools/cmd/goimports
