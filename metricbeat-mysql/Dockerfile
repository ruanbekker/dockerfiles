# Documentation: 
# https://www.elastic.co/guide/en/beats/metricbeat/current/metricbeat-module-mysql.html
FROM docker.elastic.co/beats/metricbeat:6.3.1

ENV ES_HOST    elasticsearch
ENV ES_PORT    9200
ENV ES_PROT    "http"
ENV ES_USER    elastic
ENV ES_PASS    changeme
ENV ES_INDEX  "metricbeat-%{+yyyy.MM.dd}"
ENV METRIC_PERIOD 1s
ENV MYSQL_DSN "root:pass@tcp(mysql:3306)/"
ENV LOGGING_LEVEL "warning"

USER root
COPY metricbeat.yml /usr/share/metricbeat/metricbeat.yml
# RUN chown metricbeat /usr/share/metricbeat/metricbeat.yml
# USER metricbeat
RUN chmod go-w /usr/share/metricbeat/metricbeat.yml
RUN metricbeat test config -c /usr/share/metricbeat/metricbeat.yml

ENTRYPOINT ["metricbeat"]
