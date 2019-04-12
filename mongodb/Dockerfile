FROM alpine
MAINTAINER Ruan Bekker <ruan@ruanbekker.com>

RUN apk add --no-cache mongodb mongodb-tools && mkdir -p /data/db
EXPOSE 27017

CMD ["mongod", "--syslog", "--nojournal", "--storageEngine", "wiredTiger", "--smallfiles", "--dbpath", "/data/db", "--bind_ip_all", "--port", "27017"] 
