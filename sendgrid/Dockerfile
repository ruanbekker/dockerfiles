FROM node:8-alpine

ADD package.json /app/
ADD package-lock.json /app/
ADD config.js /app/
ADD utils.js /app/
ADD server.js /app/

WORKDIR /app/

RUN npm install
EXPOSE 25

CMD ["npm", "--silent", "start"]
