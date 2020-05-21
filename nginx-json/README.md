Nginx that outputs json logs

## Usage

Run a container:

```
$ docker run -itd -p 809:80 ruanbekker/nginx-demo:json
ae9c1d183f12013c53121e3ee398ed64b264928120ebb33dfb4b964972743d85
```

Make a request:

```
$ curl -XGET -H "Host: foo.example.com" --referer "google.com" 'http://localhost:809/?foo=bar'
ok
```

View the logs:

```
$ docker logs -f ae9c1d183f12013c53121e3ee398ed64b264928120ebb33dfb4b964972743d85 | jq .
{
  "@timestamp": "2020-05-21T08:04:55+00:00",
  "host": "ae9c1d183f12",
  "server_ip": "172.17.0.2",
  "client_ip": "172.17.0.1",
  "xff": "-",
  "domain": "foo.example.com",
  "url": "/index.html",
  "referer": "google.com",
  "args": "foo=bar",
  "upstreamtime": "-",
  "responsetime": "0.000",
  "request_method": "GET",
  "status": "200",
  "size": "3",
  "request_body": "-",
  "request_length": "108",
  "protocol": "HTTP/1.1",
  "upstreamhost": "-",
  "file_dir": "/usr/share/nginx/html/index.html",
  "http_user_agent": "curl/7.64.1"
}
```
