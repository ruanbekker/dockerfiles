#!/bin/sh

CDIR=$(pwd)
cd /tmp

git clone https://github.com/kumina/postfix_exporter
cd postfix_exporter
docker run -i -v `pwd`:/postfix_exporter golang:1.12 /bin/sh << 'EOF'
set -ex

# Install prerequisites for the build process.
apt-get update -q
apt-get install -yq libsystemd-dev

cd /postfix_exporter

go get -d ./...
go build -a -tags nosystemd
strip postfix_exporter
EOF

mv postfix_exporter $CDIR/postfix_exporter
rm -rf /tmp/postfix_exporter
