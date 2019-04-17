FROM golang:alpine
ADD main.go /go/src/myapp/main.go
RUN apk add --no-cache git
WORKDIR /go/src/myapp/
RUN go get
CMD ["go", "run", "main.go"]
