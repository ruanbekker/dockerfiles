FROM alpine:latest
ARG timestamp
LABEL built_date=$timestamp
RUN apk --no-cache add curl

