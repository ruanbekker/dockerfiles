FROM alpine:3.13

ENV AWS_ACCESS_KEY_ID="${AWS_ACCESS_KEY_ID}"
ENV AWS_SECRET_ACCESS_KEY="${AWS_SECRET_ACCESS_KEY}"
ENV AWS_DEFAULT_REGION="${AWS_DEFAULT_REGION}"
ENV TERRAFORM_VERSION=0.12.15

RUN apk update \
    && apk add curl jq python bash ca-certificates git openssl ncurses unzip wget py2-pip docker \
    && pip install awscli \
    && cd /tmp \
    && wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip \
    && unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip -d /usr/bin \
    && curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl \
    && curl -LO https://raw.githubusercontent.com/ahmetb/kubectx/master/kubectx \
    && curl -LO https://raw.githubusercontent.com/ahmetb/kubectx/master/kubens \
    && chmod +x ./kube* \
    && mv ./kube* /usr/local/bin/ \
    && rm -rf /tmp/* /var/cache/apk/* rm -rf /var/tmp/*

#ENTRYPOINT [ "/bin/sh" ]
