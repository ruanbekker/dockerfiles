init:

```
docker run --rm -v "$PWD:/gitbook" -p 4000:4000 ruanbekker/gitbook:2.6.9 gitbook init
```

serve:

````
docker run --rm -v "$PWD:/gitbook" -p 4000:4000 ruanbekker/gitbook:2.6.9 gitbook serve
```
