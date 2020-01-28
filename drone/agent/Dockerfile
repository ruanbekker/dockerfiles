FROM scratch
ENV GODEBUG=netdns=go
ENV DRONE_PLATFORM=linux/amd64
ADD release/linux/arm/drone-agent /bin/

ENTRYPOINT ["/bin/drone-agent"]
