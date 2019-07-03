FROM alpine
MAINTAINER Ruan Bekker <ruan@ruanbekker.com>

WORKDIR /motd
ADD . .

RUN apk --no-cache add bash && chmod +x boot.sh

CMD ["/motd/boot.sh"]
