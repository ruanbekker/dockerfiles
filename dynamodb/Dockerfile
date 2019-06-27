FROM node:8.16.0-stretch-slim

RUN apt update && apt install build-essential python-minimal -y && mkdir /data
RUN npm install --unsafe-perm -g dynalite
RUN apt-get clean

CMD ["dynalite", "--port", "8000", "--path", "/data", "--maxItemSizeKb", "400"]
