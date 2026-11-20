# Changelog

## [0.5.1](https://github.com/eshiji/terraform-container-image/compare/0.5.0...0.5.1) (2026-11-20)


### Bug Fixes

* put go on PATH ([#15](https://github.com/eshiji/terraform-container-image/issues/15)) ([5b4c06a](https://github.com/eshiji/terraform-container-image/commit/5b4c06ac38cd5ef99a99dea93e5232cc743f9053))

## 0.5.0 - 11/20/2026
- Fix image build: install Terraform directly from releases.hashicorp.com with SHA256 verification (tfswitch v1.5.1 can no longer verify HashiCorp's PGP signature, and its ~/bin install never reached the final stage)
- Add functional test running `terraform init/validate/plan` inside the image
- Build the test image once per test run
- Read declared tool versions from the Dockerfile with nana instead of repeating them in tests; add the missing terraform check and a guard for unchecked version ARGs
- Harden the test workflow: actions pinned by SHA, read-only permissions, Go from tests/go.mod

## 0.4.3 - 09/08/2025
- Add `make help`

## 0.4.2 - 09/05/2025
- Add tag latest

## 0.4.1 - 09/05/2025
- Update .dockerignroe
- Update .github workflows names

## 0.4.0 - 09/04/2025
- Improve pre-commit

## 0.3.0 - 09/04/2025
- add PR template

## 0.2.3 - 09/04/2025
- Fix docker registry name

## 0.2.2 - 09/04/2025
- Fix push image pipeline
- Update build and test workflow pipelines

## 0.2.1 - 09/04/2025
- Fix push image pipeline

## 0.2.0 - 09/04/2025
- Update push image commands
- Update pre-commit check files to iterate over a list of files

## 0.1.1 - 09/03/2025
- Add version file check on pre-commit
- Fix pipelines runs_on

## 0.1.0 - 09/03/2025
- Add Makefile
- Update tests
- Add Changelog
- Add tag and release GH actions pipelines
- Add pre-commit hook
- Separate build and tests pipelines
- Testing multiarch build
