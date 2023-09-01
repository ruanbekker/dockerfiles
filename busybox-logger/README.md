# busybox-logger

This is a basic logger that I use to test deployment rollouts.

## Image Versions

The tags can be found on my dockerhub page:
- https://hub.docker.com/r/ruanbekker/busybox-logger

| Version  | Image                           |
|----------|---------------------------------|
| v1       | `ruanbekker/busybox-logger:v1`    |
| v2       | `ruanbekker/busybox-logger:v2`    |
| v3       | `ruanbekker/busybox-logger:v3`    |
| v4       | `ruanbekker/busybox-logger:v4`    |
| v5       | `ruanbekker/busybox-logger:v5`    |
| v6       | `ruanbekker/busybox-logger:v6`    |
| v7       | `ruanbekker/busybox-logger:v7`    |
| v8       | `ruanbekker/busybox-logger:v8`    |
| v9       | `ruanbekker/busybox-logger:v9`    |
| v10      | `ruanbekker/busybox-logger:v10`   |
| v11      | `ruanbekker/busybox-logger:v11`   |
| v12      | `ruanbekker/busybox-logger:v12`   |
| v13      | `ruanbekker/busybox-logger:v13`   |
| v14      | `ruanbekker/busybox-logger:v14`   |
| v15      | `ruanbekker/busybox-logger:v15`   |
| v16      | `ruanbekker/busybox-logger:v16`   |
| v17      | `ruanbekker/busybox-logger:v17`   |
| v18      | `ruanbekker/busybox-logger:v18`   |
| v19      | `ruanbekker/busybox-logger:v19`   |
| v20      | `ruanbekker/busybox-logger:v20`   |

## Running a container

To run the container:

```bash
docker run -it ruanbekker/busybox-logger:v12

[2023-09-01T17:45:27] level=info version=v12 message="this is the v12 container running with random id of 25387"
[2023-09-01T17:45:28] level=info version=v12 message="this is the v12 container running with random id of 26780"
[2023-09-01T17:45:30] level=info version=v12 message="this is the v12 container running with random id of 32132"
[2023-09-01T17:45:31] level=info version=v12 message="this is the v12 container running with random id of 12910"
```
