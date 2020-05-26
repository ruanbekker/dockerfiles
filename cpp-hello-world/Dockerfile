FROM ubuntu:focal

ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=Africa/Johannesburg

WORKDIR /app

RUN apt update && \
	apt install build-essential git cmake autoconf libtool pkg-config -y

ADD main.cpp .
ADD CMakeLists.txt .

RUN cmake . && make

CMD ["/app/hello"]
