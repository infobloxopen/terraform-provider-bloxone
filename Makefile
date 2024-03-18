TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=registry.terraform.io
NAMESPACE=infobloxopen
NAME=bloxone
BINARY=terraform-provider-${NAME}
VERSION=1.0.0
OS_ARCH=linux_amd64
MODULES_DIR=./modules
TERRAFORM_DOCS_IMAGE=quay.io/terraform-docs/terraform-docs:0.17.0

default: install

build:
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	cp ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4 -coverprofile cover.out

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m -coverprofile testacc-cover.out

gen: modules-docs
	go generate

modules-docs: $(MODULES_DIR)/*
	@for d in $^ ; do \
		echo "Generating documentation for module $$d" ; \
		docker run --rm --volume "./$$d:/$$d" $(TERRAFORM_DOCS_IMAGE) markdown "/$$d" ; \
	done

fmt:
	go fmt ./...

.PHONY: default test testacc gen fmt
