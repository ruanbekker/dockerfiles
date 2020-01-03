FROM golang:alpine as builder
ENV GOBIN ${GOPATH}/bin
RUN apk add git
COPY app.go /go/app.go
WORKDIR /go
RUN go get && go build -tags=netgo -o /go/app

FROM scratch
COPY --from=builder /go/app /app
CMD ["/app"]
