# basic-api

Basic web application for testing

## Usage

Using environment variables to control the output:

```bash
docker run -it -p 8080:8080 -e API_VERSION=v2 ruanbekker/basic-api:latest
```

Using tags for the version:

```bash
docker run -it -p 8080:8080 ruanbekker/basic-api:v1
```

Available tags:

- `ruanbekker/basic-api:v1`
- `ruanbekker/basic-api:v2`
- `ruanbekker/basic-api:v3`
