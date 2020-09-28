FROM python:3.7-alpine

ENV ANSIBLE_VERSION=2.10.0

RUN apk add --no-cache --virtual=.build-deps libffi-dev openssl-dev build-base \
    && apk add --no-cache --virtual=.run-deps openssh-client ca-certificates openssl \
    && pip install --no-cache-dir cffi==1.14.3 ansible==${ANSIBLE_VERSION} \
    && apk del .build-deps \
    && ln -s /usr/local/bin/python /usr/bin/python
    
CMD ["ansible-playbook", "--version"]
