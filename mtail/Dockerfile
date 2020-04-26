FROM golang:alpine

EXPOSE 9101

ENV GOPATH     /go
ENV PATH       /go/bin:$PATH

RUN apk add --update musl go git bash make && \
	mkdir /mtail  && \
	go get github.com/google/mtail/cmd/mtail && \
	apk del --purge musl go git

COPY nginx.mtail /opt/mtail/progs/nginx.mtail

ENTRYPOINT ["/go/bin/mtail", "--address", "0.0.0.0", "--port", "9101", "--emit_prog_label=false", "--alsologtostderr", "--progs", "/opt/mtail/progs", "--logs"]
