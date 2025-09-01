GO_MOD_NAME=github.com/eshiji/terraform-docker
TERRAFORM_IMAGE=eshijii/terraform-tools

.PHONY: go_init
go_init:
	cd tests && \
	go mod init ${GO_MOD_NAME}

.PHONY: test
test:
	cd tests && \
	go mod tidy && \
	go test

.PHONY: build
build:
	docker build -t ${TERRAFORM_IMAGE} .
