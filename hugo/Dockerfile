FROM alpine

LABEL description="Hugo Static Site Generator for Docker"
LABEL maintainer="Ruan Bekker <ruan@ruanbekker.com>"

ENV HUGO_VERSION=0.69.0
ENV HUGO_HOME=/usr/local/hugo
ENV PROJECT_PATH=/app
ENV PATH=${PATH}:${HUGO_HOME}/bin

ADD https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.tar.gz /tmp
RUN apk add --no-cache ca-certificates git \
    && tar -xf /tmp/hugo_${HUGO_VERSION}_Linux-64bit.tar.gz -C /tmp \
    && mkdir -p ${PROJECT_PATH} \
    && mkdir -p ${HUGO_HOME}/bin \
    && mv /tmp/hugo ${HUGO_HOME}/bin/hugo \
    && rm -rf /tmp/hugo_* \
    && rm -rf /tmp/*.md
