# echo container

Container for demonstrations, nothing fancy

## Build

Multi-arch build:

```bash
$ docker buildx build --platform "linux/amd64,linux/arm64" -f Dockerfile -t ruanbekker/containers:echo  --push .
```

## Usage

```bash
$ docker run -it ruanbekker/containers:echo
2022-09-15T09:09:49Z : ping from 62be534d632a : iteration 1
2022-09-15T09:09:54Z : ping from 62be534d632a : iteration 2
```
