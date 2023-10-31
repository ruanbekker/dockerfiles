FROM ubuntu:22.04

WORKDIR /tmp

RUN apt update && \
    # Install build dependencies
    apt install -y build-essential autoconf automake libtool pkg-config git libssl-dev libc-ares-dev && \

    # Build and install curl
    git clone --depth 1 https://github.com/curl/curl.git && \
    cd curl && \
    autoreconf -fi && \
    ./configure --enable-ares --with-openssl && \
    make && \
    make install && \
    ldconfig && \

    # Remove build dependencies and cleanup
    apt purge -y build-essential autoconf automake libtool pkg-config git libssl-dev libc-ares-dev && \
    apt clean autoclean && \
    apt autoremove -y && \
    rm -rf /var/lib/{apt,dpkg,cache,log}/ && \
    rm -rf /tmp/*

ENTRYPOINT ["/usr/local/bin/curl"]

# Usage:
# curl --dns-servers 10.10.200.2 --location 'http://some-endpoint.lan' 
