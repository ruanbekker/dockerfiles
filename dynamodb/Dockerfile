FROM frolvlad/alpine-oraclejdk8

RUN apk update \
    && apk add python openssl \
    && apk add py2-pip \
    && pip install awscli \
    && mkdir -p /opt/ddb_local \
    && cd /opt/ddb_local \
    && wget http://dynamodb-local.s3-website-us-west-2.amazonaws.com/dynamodb_local_latest.tar.gz \
    && tar -xf dynamodb_local_latest.tar.gz \
    && echo 'java -Djava.library.path=/opt/ddb_local/DynamoDBLocal_lib -jar /opt/ddb_local/DynamoDBLocal.jar - inMemory -sharedDb' > /opt/ddb.sh \
    && chmod +x /opt/ddb.sh

ENTRYPOINT ["sh"]
CMD ["/opt/ddb.sh"]

EXPOSE 8000
