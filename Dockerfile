FROM ubuntu:20.04 AS bin

LABEL maintainer="eshijii@gmail.com"

ARG GOLANG_VERSION="1.18.2"
ARG ARCH="amd64"
ARG TERRAFORM_VERSION="1.2.1"

ARG TFSWITCH_DOWNLOAD_URL="https://raw.githubusercontent.com/warrensbox/terraform-switcher/release/install.sh"
ARG TERRAGRUNT_URL="https://github.com/gruntwork-io/terragrunt/releases/download/v0.36.10/terragrunt_linux_amd64"
ARG GOLANG_URL="https://golang.org/dl/go${GOLANG_VERSION}.linux-${ARCH}.tar.gz"

RUN apt-get update && apt-get install -y \
    build-essential \
    curl \
    wget \
    && rm -rf /var/lib/apt/lists/*

# tfswitch
RUN curl -L $TFSWITCH_DOWNLOAD_URL | bash && \
    tfswitch $TERRAFORM_VERSION

# terragrunt
RUN wget -cq $TERRAGRUNT_URL -O /usr/local/bin/terragrunt && \
    chmod u+x /usr/local/bin/terragrunt

# golang
RUN wget -cq $GOLANG_URL && \
    tar -xf "go${GOLANG_VERSION}.linux-${ARCH}.tar.gz" && \
    mv -v go /usr/local && \
    rm -rf "go${GOLANG_VERSION}.linux-${ARCH}.tar.gz"

ENV GOPATH=$HOME/go
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin