FROM public.ecr.aws/ubuntu/ubuntu:22.04

ENV SHARED_CONFIG_DIR=/opt/coder

RUN apt update && apt install supervisor -y

RUN useradd coder
RUN mkdir -p /opt/coder
RUN chown coder /opt/coder
USER coder

RUN mkdir -p $SHARED_CONFIG_DIR/etc/supervisor $SHARED_CONFIG_DIR/logs
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

COPY app-colors/colors /usr/bin/colors
COPY app-messages/messages /usr/bin/messages

WORKDIR /opt/coder

CMD ["/usr/bin/supervisord"]
