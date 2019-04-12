FROM alpine
MAINTAINER Ruan Bekker <ruan@ruanbekker.com>

RUN apk add --no-cache mongodb mongodb-tools && mkdir -p /data/db
EXPOSE 27017

CMD ["mongod", "--syslog", "--nojournal", "--smallfiles", "--dbpath", "/data/db", "--bind_ip_all"] 
