## Whoami

Golang app that returns the hostname

### Usage:

Run a container:

```
$ docker run -p 8000:8000 ruanbekker/whoami:me                                                                                                                     2 ↵
Server listening on port 8000
```

Make a request:

```
$ curl http://localhost:8000/                                                                                                                             130 ↵
Hostname: ea75b484dbaf
```

About me:

```
$ curl -s http://localhost:8000/about | jq .
{
  "name": "Ruan",
  "twitter": "@ruanbekker",
  "contact": "https://ruan.dev",
  "blogs": [
    "https://blog.ruanbekker.com",
    "https://sysadmins.co.za",
    "https://ruan.dev/blog"
  ]
}
```

### Resources:

- https://gowebexamples.com/http-server/
- https://gobyexample.com/environment-variables
- https://www.golang-book.com/books/intro/4
- https://flaviocopes.com/go-web-server/
- https://www.alexedwards.net/blog/golang-response-snippets
