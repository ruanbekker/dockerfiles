```
$ mkdir code 
$ mkdir data

$ docker run --rm --name vscode \
  -it -p 8443:8443 -p 8888:8888 \
  -v $(pwd)/data:/data -v $(pwd)/code:/code \
  ruanbekker/vscode:python-3.7
```

Docker

```
$ docker run --privileged -it -p 8443:8443 ruanbekker/vscode:docker
```
