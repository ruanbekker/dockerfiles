FROM busybox

ARG version=v1
ENV version=$version

COPY run.sh /run.sh
RUN chmod +x /run.sh

CMD ["/run.sh"]
