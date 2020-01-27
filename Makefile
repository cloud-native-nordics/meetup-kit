PROJECT=github.com/cloud-native-nordics/meetup-kit
GO_VERSION=1.13.1
BINARIES=meetup-kit
CACHE_DIR = $(shell pwd)/bin/cache
BOUNDING_API_DIRS = ${PROJECT}/pkg/apis
API_DIRS = ${PROJECT}/pkg/apis/meetops,${PROJECT}/pkg/apis/meetops/v1alpha1

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

autogen:
	$(MAKE) shell COMMAND="make dockerized-autogen"

dockerized-autogen: /go/bin/deepcopy-gen /go/bin/defaulter-gen /go/bin/conversion-gen /go/bin/openapi-gen
	# Let the boilerplate be empty
	touch /tmp/boilerplate
	/go/bin/deepcopy-gen \
		--input-dirs ${API_DIRS} \
		--bounding-dirs ${BOUNDING_API_DIRS} \
		-O zz_generated.deepcopy \
		-h /tmp/boilerplate 

	/go/bin/defaulter-gen \
		--input-dirs ${API_DIRS} \
		-O zz_generated.defaults \
		-h /tmp/boilerplate

	/go/bin/conversion-gen \
		--input-dirs ${API_DIRS} \
		-O zz_generated.conversion \
		-h /tmp/boilerplate
	
	/go/bin/openapi-gen \
		--input-dirs ${API_DIRS} \
		--output-package ${PROJECT}/api/openapi \
		--report-filename api/openapi/violations.txt \
		-h /tmp/boilerplate

/go/bin/deepcopy-gen /go/bin/defaulter-gen /go/bin/conversion-gen: /go/bin/%:
	go get k8s.io/code-generator/cmd/$*

/go/bin/openapi-gen:
	go get k8s.io/kube-openapi/cmd/openapi-gen
