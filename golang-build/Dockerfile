FROM alpine
ENV GO_VERSION 1.15.2
RUN apk add --no-cache ca-certificates gcc musl-dev git && echo 'hosts: files dns' > /etc/nsswitch.conf
RUN apk add --no-cache --virtual .build-deps bash gcc musl-dev openssl go
ENV GOROOT_BOOTSTRAP="$(go env GOROOT)"
#ENV GOOS="$(go env GOOS)"
#ENV GOARCH="$(go env GOARCH)"
#ENV GOHOSTOS="$(go env GOHOSTOS)"
#ENV GOHOSTARCH="$(go env GOHOSTARCH)"
RUN wget -O go.tgz "https://golang.org/dl/go$GO_VERSION.src.tar.gz" && \
    tar -C /usr/local -xzf go.tgz && \
    rm go.tgz && \
    cd /usr/local/go/src && \
    ./make.bash && \
    rm -rf /usr/local/go/pkg/bootstrap /usr/local/go/pkg/obj && \
    apk del .build-deps

ENV PATH="/usr/local/go/bin:$PATH"
ENV GOPATH=/go
ENV export PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin"
WORKDIR $GOPATH

CMD ["/bin/sh"]
