## Dockerhub

The dockerhub image can be found at:
- https://hub.docker.com/r/ruanbekker/awscli

## Usage

Example

```
$ docker run -it --rm \
  -e AWS_ACCESS_KEY_ID=123 \
  -e AWS_SECRET_ACCESS_KEY=xyz \
  -e AWS_DEFAULT_REGION: eu-west-1 \
  ruanbekker/awscli s3 ls /
```

Or using functions:

```
aws(){
  docker run --rm -it -v ~/.aws/credentials:/root/.aws/credentials:ro ruanbekker/awscli "$@"
}
```

Or alias:

```
alias aws="docker run --rm -it -v ~/.aws/credentials:/root/.aws/credentials:ro ruanbekker/awscli"
```

## Dockerhub Images

- [ruanbekker/awscli:latest](https://hub.docker.com/r/ruanbekker/awscli/tags?page=1&ordering=last_updated)
- [ruanbekker/awscli:1.20.33](https://hub.docker.com/r/ruanbekker/awscli/tags?page=1&ordering=last_updated)
