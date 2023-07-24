FROM golang:1.16-buster AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN go build -o /tmp/app

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=builder /tmp/app /app
ENTRYPOINT ["/app"]
