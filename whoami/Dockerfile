FROM golang:alpine AS builder

WORKDIR /go/src/myapp/

COPY app.go .

ARG CGO_ENABLED=0
RUN go build -o /go/app

FROM scratch
COPY --from=builder /go/app /go/app

CMD ["/go/app"]
