FROM alpine AS builder
ENV DOCKER_VERSION 19.03.10
ENV DOCKER_ARCH x86_64

RUN apk add --no-cache wget openssl \
    && wget https://download.docker.com/linux/static/stable/${DOCKER_ARCH}/docker-${DOCKER_VERSION}.tgz \
    && tar -xf docker-${DOCKER_VERSION}.tgz && mv docker/docker /opt/docker

FROM busybox
COPY --from=builder /opt/docker /bin/docker

CMD ["/bin/sh"]
