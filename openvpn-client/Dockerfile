FROM alpine:3.7

RUN apk --update add curl openvpn

HEALTHCHECK --interval=60s --timeout=15s --start-period=120s \
  CMD curl 'https://api.ipify.org'

VOLUME [ "/vpn/config" ]
CMD [ "openvpn", "--config", "/vpn/config/config.ovpn" ]
