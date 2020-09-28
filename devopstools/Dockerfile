FROM python:3.7-alpine

ENV TERRAFORM_VERSION=0.12.12
ENV AWS_ACCESS_KEY_ID="${AWS_ACCESS_KEY_ID}"
ENV AWS_SECRET_ACCESS_KEY="${AWS_SECRET_ACCESS_KEY}"
ENV AWS_DEFAULT_REGION="${AWS_DEFAULT_REGION}"

RUN apk add --no-cache --virtual=.build-deps libffi-dev openssl-dev build-base \
    && apk add --no-cache --virtual=.run-deps openssh-client ca-certificates openssl \
    && pip install --no-cache-dir cffi==1.13.2 ansible==2.9.2 \
    && apk del .build-deps \
    && ln -s /usr/local/bin/python /usr/bin/python

RUN apk add --no-cache curl jq python bash ca-certificates git openssl ncurses unzip wget py2-pip \
    && pip install awscli \
    && cd /tmp  \
    && wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip \
    && unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip -d /usr/bin  \
    && curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl \
    && curl -LO https://raw.githubusercontent.com/ahmetb/kubectx/master/kubectx \
    && curl -LO https://raw.githubusercontent.com/ahmetb/kubectx/master/kubens \
    && chmod +x ./kube* \
    && mv ./kube* /usr/local/bin/ \
    && rm -rf /tmp/* \
    && rm -rf /var/cache/apk/* \
    && rm -rf /var/tmp/*
ENV PS1 "\h > "
CMD ["/bin/bash"]
