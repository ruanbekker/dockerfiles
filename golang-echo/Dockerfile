FROM golang:alpine AS builder

WORKDIR $GOPATH/src/echo/
COPY go.mod .
COPY main.go .
RUN go build -o /go/app

FROM scratch

COPY --from=builder /go/app /go/app

CMD ["/go/app"]
