FROM python:3.8-alpine

ENV ANSIBLE_VERSION 4.4.0

RUN apk add --no-cache --virtual .build-deps musl-dev python3-dev libffi-dev openssl-dev cargo gcc \
    && pip install ansible==$ANSIBLE_VERSION \
    && apk del .build-deps musl-dev python3-dev libffi-dev openssl-dev cargo gcc

CMD ["sh"]
