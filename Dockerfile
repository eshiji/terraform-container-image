FROM ubuntu:18.04 AS bin

LABEL maintainer="eshijii@gmail.com"

ENV TFSWITCH_DOWNLOAD_URL="https://raw.githubusercontent.com/warrensbox/terraform-switcher/release/install.sh"
ENV TERRAGRUNT_URL="https://github.com/gruntwork-io/terragrunt/releases/download/v0.36.10/terragrunt_linux_amd64"

ENV GO_VERSION="1.18.1"
ENV GO_ARCH="amd64"

RUN apt-get update && apt-get install -y \
    curl \
    wget

# tfswitch
RUN curl -L $TFSWITCH_DOWNLOAD_URL | bash

# terragrunt
RUN wget -cq $TERRAGRUNT_URL -O /usr/local/bin/terragrunt && \
    chmod u+x /usr/local/bin/terragrunt

FROM golang:1.18.1-stretch

ENV TF_VERSION="1.1.9"

COPY --from=bin /usr/local/bin/tfswitch /usr/local/bin/tfswitch
COPY --from=bin /usr/local/bin/terragrunt /usr/local/bin/terragrunt

RUN export TF_VERSION=$TF_VERSION && \
    tfswitch