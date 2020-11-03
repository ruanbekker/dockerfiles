FROM alpine AS builder
ENV WAYPOINT_VERSION 0.1.4
ENV WAYPOINT_OS linux
ENV WAYPOINT_ARCH amd64

RUN apk add --no-cache wget openssl unzip \
    && wget https://releases.hashicorp.com/waypoint/${WAYPOINT_VERSION}/waypoint_${WAYPOINT_VERSION}_${WAYPOINT_OS}_${WAYPOINT_ARCH}.zip \
    && unzip waypoint_${WAYPOINT_VERSION}_${WAYPOINT_OS}_${WAYPOINT_ARCH}.zip \
    && mv waypoint /opt/waypoint \
    && chmod +x /opt/waypoint

FROM busybox
COPY --from=builder /opt/waypoint /bin/waypoint

CMD ["/bin/sh"]
