ARG VERSION
# https://github.com/elastic/logstash-docker
FROM docker.elastic.co/logstash/logstash-oss:${VERSION}
RUN logstash-plugin install logstash-filter-json
