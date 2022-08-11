FROM golang:alpine AS builder
WORKDIR /go/src/hello
RUN apk add --no-cache gcc libc-dev
ADD app.go /go/src/hello/app.go
RUN GOOS=linux GOARCH=amd64 go build -tags=netgo app.go

FROM busybox
COPY --from=builder /go/src/hello/app /app
CMD ["/app"]
