# terraform-docker
Docker image with tools to build and test terraform modules.

The image contains:
 - Golang (1.18.2) to run terratest tests
 - [Terragrunt](https://terragrunt.gruntwork.io/)
 - Terraform 1.2.1
 - [tfswitch](https://tfswitch.warrensbox.com)

## Testing
For running container binaries tests terratest tool is used.

To run tests:
```shell
make test
```
