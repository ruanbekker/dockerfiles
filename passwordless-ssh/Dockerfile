FROM alpine

RUN apk --no-cache add \
    bash \
    curl \
    openssh \
    openssl \
    bind-tools \
  && ssh-keygen -f /etc/ssh/ssh_host_rsa_key -N '' -t rsa

COPY sshd_config /etc/ssh/sshd_config
COPY boot.sh /boot.sh
COPY motd-generator.sh /motd-generator.sh
RUN chmod +x /boot.sh && addgroup ssh-users

CMD ["/boot.sh"]
