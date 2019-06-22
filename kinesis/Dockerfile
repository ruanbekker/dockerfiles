FROM node:8.16.0-stretch-slim

RUN apt update && apt install build-essential python-minimal -y
RUN npm install --unsafe-perm -g kinesalite
RUN apt-get clean

CMD ["kinesalite", "--port", "4567", "--createStreaMs", "5"]
