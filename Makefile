GO_MOD_NAME=github.com/eshiji/terraform-docker
TERRAFORM_IMAGE=eshijii/terraform-tools

.PHONY: go_init
go_init: ## Initialize go module for terratests
	cd tests && \
	go mod init ${GO_MOD_NAME}

.PHONY: test
test: ## Execute terratests
	cd tests && \
	go mod tidy && \
	go test

.PHONY: build
build: ## Build Docker Image
	docker build -t ${TERRAFORM_IMAGE} .

.PHONY: help
help: ## Help Menu
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-40s\033[0m \t\t%s\n", $$1, $$2}'
