Credit to: https://github.com/FGRibreau/smtp-to-sendgrid-gateway

## Run

```shell
docker run -it -p 25:25 -e SENDGRID_API=XXXXXXX fgribreau/smtp-to-sendgrid-gateway
```

### Configuration - environment variables

- `SENDGRID_API` (required): sendgrid API token
- `SMTP_PORT` (optional) Port to listen (default: 25)

