# terraform-docker
Docker image with tools to build and test terraform modules.

The image contains:
 - Golang (1.18.2) to run terratest tests
 - [Terragrunt](https://terragrunt.gruntwork.io/)
 - Terraform 1.2.1 
 - [tfswitch](https://tfswitch.warrensbox.com) 

## Testing 

To run tests:
```shell
cd test
go test
```
ARG GOLANG_VERSION="1.18.2"
ARG ARCH="amd64"
ARG TERRAFORM_VERSION="1.2.1"
