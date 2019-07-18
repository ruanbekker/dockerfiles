FROM node:8.16.0-stretch

ENV DYNAMO_ENDPOINT=${DYNAMO_ENDPOINT}

RUN npm install dynamodb-admin -g

CMD ["dynamodb-admin", "--port", "8000"]
