FROM golang:alpine AS build
RUN apk add --no-cache git
RUN go get github.com/jsha/minica

FROM busybox
COPY --from=build /go/bin/minica /usr/bin/minica
RUN mkdir /output
WORKDIR /output
ENTRYPOINT ["minica"]
