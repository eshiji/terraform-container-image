FROM --platform=$BUILDPLATFORM golang:alpine AS go_install
# BuildKit automatically sets BUILDPLATFORM (e.g., linux/amd64) for the build stage's base image.
# The `golang:alpine` base image is multi-arch, so it will match the builder's architecture.

# Set GOOS and GOARCH to match the final image's TARGETPLATFORM.
# BuildKit automatically sets TARGETOS and TARGETARCH.
ARG TARGETOS TARGETARCH
ARG GOLANG_VERSION="1.25.0"

WORKDIR /go

RUN case "$TARGETARCH" in \
        "amd64") ARCH=amd64;; \
        "arm64") ARCH=arm64;; \
        "darwin64") ARCH=darwin64;; \
        *) echo "unsupported architecture: $TARGETARCH"; exit 1;; \
    esac; \
    wget -cq "https://golang.org/dl/go${GOLANG_VERSION}.${TARGETOS}-${ARCH}.tar.gz" && \
    tar -xf "go${GOLANG_VERSION}.${TARGETOS}-${ARCH}.tar.gz"

############ END ############

FROM --platform=$BUILDPLATFORM ubuntu:24.04 AS bin_install

LABEL maintainer="eshijii@gmail.com"

ARG TARGETOS TARGETARCH

ARG TERRAFORM_VERSION="1.2.1"
ARG TERRAGRUNT_VERSION="v0.86.2"
ARG TERRAFORM_DOCS_VERSION="v0.20.0"
ARG TFSEC_VERSION="v1.28.14"
ARG TFSWITCH_VERSION="v1.5.1"

ARG TFSWITCH_URL="https://raw.githubusercontent.com/warrensbox/terraform-switcher/release/install.sh"
ARG TERRAGRUNT_URL="https://github.com/gruntwork-io/terragrunt/releases/download/${TERRAGRUNT_VERSION}/terragrunt_${TARGETOS}_${TARGETARCH}"
ARG TERRAFORM_DOCS_URL="https://terraform-docs.io/dl/${TERRAFORM_DOCS_VERSION}/terraform-docs-${TERRAFORM_DOCS_VERSION}-${TARGETOS}-${TARGETARCH}.tar.gz"
ARG TFSEC_URL="https://github.com/aquasecurity/tfsec/releases/download/${TFSEC_VERSION}/tfsec-${TARGETOS}-${TARGETARCH}"

RUN apt-get update && apt-get install -y \
    build-essential \
    curl \
    wget \
    && rm -rf /var/lib/apt/lists/*

# tfswitch
RUN curl -L $TFSWITCH_URL | bash -s -- $TFSWITCH_VERSION && \
    tfswitch $TERRAFORM_VERSION

# terragrunt
RUN wget -cq $TERRAGRUNT_URL -O /usr/local/bin/terragrunt && \
    chmod u+x /usr/local/bin/terragrunt

# terraform-docs
RUN wget -cqO /tmp/terraform-docs.tar.gz $TERRAFORM_DOCS_URL && \
    tar -xzf /tmp/terraform-docs.tar.gz && \
    mv terraform-docs /usr/local/bin/terraform-docs

# tfsec
RUN wget -cq $TFSEC_URL -O /usr/local/bin/tfsec && \
    chmod u+x /usr/local/bin/tfsec

############ END ############

FROM --platform=$BUILDPLATFORM ubuntu:24.04
COPY --from=go_install /go /usr/local/bin/
COPY --from=bin_install /usr/local/bin /usr/local/bin/

ENV GOPATH=$HOME/go
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
