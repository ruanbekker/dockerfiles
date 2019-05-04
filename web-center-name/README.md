## Docker Hub

- [ruanbekker/web-center-name](https://hub.docker.com/r/ruanbekker/web-center-name)

## Usage

Staging:

```
docker run -it -e APP_ENVIRONMENT=Staging -p 81:5000 ruanbekker/web-center-name
```

Production:

```
docker run -it -e APP_ENVIRONMENT=Production -p 80:5000 ruanbekker/web-center-name
```

## Screenshots

Staging:

<img width="1259" alt="image" src="https://user-images.githubusercontent.com/567298/57185377-73fe3700-6eca-11e9-91d3-953e754cbde9.png">

Production:

<img width="1270" alt="image" src="https://user-images.githubusercontent.com/567298/57185383-8d06e800-6eca-11e9-8cff-c3a665f9f377.png">
