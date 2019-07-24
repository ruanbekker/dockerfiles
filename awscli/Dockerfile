FROM alpine:latest

ENV AWS_ACCESS_KEY_ID="${AWS_ACCESS_KEY_ID}"
ENV AWS_SECRET_ACCESS_KEY="${AWS_SECRET_ACCESS_KEY}"
ENV AWS_DEFAULT_REGION="${AWS_DEFAULT_REGION}"

RUN apk --no-cache add \
	ca-certificates \
	python \
	py2-pip \
	&& pip install awscli
	
ENTRYPOINT [ "aws" ]
