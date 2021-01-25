FROM alpine

RUN apk --no-cache add wget iptables lxc tar curl

RUN wget https://download.docker.com/linux/static/stable/x86_64/docker-19.03.10.tgz \
    && tar -xf docker-19.03.10.tgz \
    && rm -rf docker-19.03.10.tgz \
    && mv docker/* /usr/local/bin/ \
    && rm -rf docker

ADD entrypoint.sh /bin/entrypoint.sh
RUN chmod +x /bin/entrypoint.sh

ENTRYPOINT ["/bin/entrypoint.sh"]
CMD ["/bin/sh"]
