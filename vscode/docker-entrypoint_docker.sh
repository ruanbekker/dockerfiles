#!/bin/bash
set -e

if [ $# -eq 0 ]
  then
    screen -S dockerd -m -d bash -c "dockerd"
    /usr/local/bin/code-server --allow-http --data-dir /data /code
  else
    exec "$@"
fi
