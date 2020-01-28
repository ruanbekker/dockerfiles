#!/usr/bin/env bash

DRONE_BUILD_NUMBER=1.2.3
wget https://github.com/drone/drone/archive/v${DRONE_BUILD_NUMBER}.zip
unzip v$DRONE_BUILD_NUMBER.zip
cd drone-$DRONE_BUILD_NUMBER/cmd/drone-server/
go get
GOOS=linux GOARCH=amd64 go build -ldflags '-X github.com/drone/drone/version.VersionDev=build.'${DRONE_BUILD_NUMBER} -o release/linux/amd64/drone-server github.com/drone/drone/cmd/drone-server
mv release ../../../
cd ../../..
