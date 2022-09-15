# containers

Usefule container images for debug scenarios and ci pipelines

## Image Registry

Container images are hosted on Docker Hub and can be viewed:
- [hub.docker.com/r/ruanbekker/containers](https://hub.docker.com/r/ruanbekker/containers)

## Buildx

Docker Buildx is a CLI plugin that extends the docker command with the full support of the features provided by Moby BuildKit builder toolkit.

```bash
docker buildx create --name multi-arch --platform "linux/arm64,linux/amd64,linux/arm/v7" --driver "docker-container"
docker buildx use multi-arch
```

Building:

```bash
docker buildx build \
  --platform "linux/amd64,linux/arm64,linux/arm/v7" \
  -f curl.Dockerfile \
  --build-arg timestamp=$(date +%F) \
  -t ruanbekker/containers:curl \
  --push .
```

## Images

The following tags are available

| Name   | Container Image               | Docker Hub                                                                                           |
| ------ |:-----------------------------:| ----------------------------------------------------------------------------------------------------:|
| curl   | `ruanbekker/containers:curl`  | [ruanbekker/containers:curl](https://hub.docker.com/r/ruanbekker/containers/tags?page=1&name=curl)   |
| echo   | `ruanbekker/containers:echo`  | [ruanbekker/containers:echo](https://hub.docker.com/r/ruanbekker/containers/tags?page=1&name=echo)   |
| mysql  | `ruanbekker/containers:mysql` | [ruanbekker/containers:mysql](https://hub.docker.com/r/ruanbekker/containers/tags?page=1&name=mysql) |
| nginx  | `ruanbekker/containers:nginx` | [ruanbekker/containers:nginx](https://hub.docker.com/r/ruanbekker/containers/tags?page=1&name=nginx) |


