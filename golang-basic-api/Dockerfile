FROM golang:1.17 as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

FROM gcr.io/distroless/base-debian10

COPY --from=builder /app/myapp /myapp

ENV API_VERSION=v3

CMD ["/myapp"]
