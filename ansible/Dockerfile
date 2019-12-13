FROM python:3.7-alpine

RUN apk add --no-cache --virtual=.build-deps libffi-dev openssl-dev build-base \
    && apk add --no-cache --virtual=.run-deps openssh-client ca-certificates openssl \
    && pip install --no-cache-dir cffi==1.13.2 ansible==2.9.2 \
    && apk del .build-deps \
    && ln -s /usr/local/bin/python /usr/bin/python
    
CMD ["ansible-playbook", "--version"]
