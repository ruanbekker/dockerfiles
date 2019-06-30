FROM docker.elastic.co/beats/filebeat:6.3.1

ENV ES_HOST    elasticsearch
ENV ES_PORT    9200
ENV ES_PROT    "http"
ENV ES_USER    elastic
ENV ES_PASS    changeme
ENV ES_INDEX  "filebeat-%{+yyyy.MM.dd}"
ENV LOGGING_LEVEL "warning"

RUN filebeat modules enable mysql
USER root

COPY filebeat.yml /usr/share/filebeat/filebeat.yml
COPY mysql.yml  /usr/share/filebeat/modules.d/mysql.yml
RUN chmod go-w /usr/share/filebeat/filebeat.yml

RUN filebeat test config -c /usr/share/filebeat/filebeat.yml

ENTRYPOINT ["filebeat"]
