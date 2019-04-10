FROM openjdk:11.0-jre-slim

WORKDIR /usr/local/dynamodb
ADD https://s3-us-west-2.amazonaws.com/dynamodb-local/dynamodb_local_latest.tar.gz .
RUN tar -xf dynamodb_local_latest.tar.gz \
    && rm -rf  dynamodb_local_latest.tar.gz

EXPOSE 8000
ENTRYPOINT ["/usr/bin/java", "-Djava.library.path=.", "-jar", "DynamoDBLocal.jar", "-inMemory", "-sharedDb"]
CMD ["-port", "8000"]
