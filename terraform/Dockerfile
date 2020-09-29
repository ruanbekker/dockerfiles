FROM alpine:3.8

ENV TERRAFORM_VERSION=0.13.3
ENV AWS_ACCESS_KEY_ID="${AWS_ACCESS_KEY_ID}"
ENV AWS_SECRET_ACCESS_KEY="${AWS_SECRET_ACCESS_KEY}"
ENV AWS_DEFAULT_REGION="${AWS_DEFAULT_REGION}"

RUN apk update && \
    apk add curl jq python bash ca-certificates git openssl unzip wget py2-pip && \
    pip install awscli && \
    cd /tmp && \
    wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip -d /usr/bin && \
    rm -rf /tmp/* && \
    rm -rf /var/cache/apk/* && \
    rm -rf /var/tmp/*
  	

#ENTRYPOINT ["/usr/bin/terraform"]
#CMD ["--help"]
