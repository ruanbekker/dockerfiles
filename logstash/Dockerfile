ENV VERSION
# https://github.com/elastic/logstash-docker
FROM docker.elastic.co/logstash/logstash-oss:${VERSION:-6.3.1}
RUN logstash-plugin install logstash-filter-json
ADD logstash.conf /usr/share/logstash/pipeline/logstash.conf
CMD ["logstash", "-f", "/usr/share/logstash/pipeline/logstash.conf"]
