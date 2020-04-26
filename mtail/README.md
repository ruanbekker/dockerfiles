```
$ docker run -itd --name nginx-log-exporter \
  -p 9101:9101 \
  -v /var/log/nginx:/var/log/nginx \
  ruanbekker/mtail:nginx /var/log/nginx/access_prometheus.log
```
