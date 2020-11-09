FROM python:3.7-alpine

ENV TERRAFORM_VERSION=0.13.3
ENV ANSIBLE_VERSION=2.10.0
ENV AWS_ACCESS_KEY_ID="${AWS_ACCESS_KEY_ID}"
ENV AWS_SECRET_ACCESS_KEY="${AWS_SECRET_ACCESS_KEY}"
ENV AWS_DEFAULT_REGION="${AWS_DEFAULT_REGION}"

RUN apk add --no-cache --virtual=.build-deps libffi-dev openssl-dev build-base \
    && apk add --no-cache --virtual=.run-deps openssh-client ca-certificates openssl \
    && pip install --no-cache-dir cffi==1.14.3 ansible==${ANSIBLE_VERSION} \
    && apk del .build-deps \
    && ln -s /usr/local/bin/python /usr/bin/python

RUN apk add --no-cache curl jq bash ca-certificates git openssl ncurses unzip wget iptables git mysql-client \
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
    
RUN wget https://download.docker.com/linux/static/stable/x86_64/docker-19.03.10.tgz \
    && tar -xf docker-19.03.10.tgz \
    && rm -rf docker-19.03.10.tgz \
    && mv docker/* /usr/local/bin/ \
    && rm -rf docker

# ansible galaxy
RUN pip install --upgrade --user openshift && ansible-galaxy collection install community.kubernetes

ENV WAYPOINT_VERSION 0.1.4
ENV WAYPOINT_OS linux
ENV WAYPOINT_ARCH amd64

RUN wget https://releases.hashicorp.com/waypoint/${WAYPOINT_VERSION}/waypoint_${WAYPOINT_VERSION}_${WAYPOINT_OS}_${WAYPOINT_ARCH}.zip \
    && unzip waypoint_${WAYPOINT_VERSION}_${WAYPOINT_OS}_${WAYPOINT_ARCH}.zip \
    && mv waypoint /usr/local/bin/waypoint \
    && chmod +x /usr/local/bin/waypoint \
    && rm -rf waypoint_${WAYPOINT_VERSION}_${WAYPOINT_OS}_${WAYPOINT_ARCH}.zip

RUN echo Terraform: ${TERRAFORM_VERSION} > /root/devopstools_info.txt &&  \
    echo Ansible: ${ANSIBLE_VERSION} >> /root/devopstools_info.txt && \
    echo Kubectl: $(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt) >> /root/devopstools_info.txt

ENV PS1 "\h > "
CMD ["/bin/bash"]

