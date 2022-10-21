FROM golang:alpine AS builder
WORKDIR /go/src/hello
RUN apk add --no-cache gcc libc-dev
ADD app.go /go/src/hello/app.go
RUN GOOS=linux go build -tags=netgo app.go
#RUN GOOS=linux GOARCH=arm64 go build -tags=netgo app.go

FROM scratch
COPY --from=builder /go/src/hello/app /app
CMD ["/app"]
