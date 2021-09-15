FROM golang:alpine AS builder
WORKDIR /go/src/hello
RUN apk add --no-cache gcc libc-dev
ADD go.* /go/src/hello/
ADD app.go .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -tags=netgo app.go

FROM scratch
COPY --from=builder /go/src/hello/app /app
CMD ["/app"]
