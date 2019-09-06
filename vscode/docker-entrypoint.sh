#!/bin/bash
set -e

if [ $# -eq 0 ]
  then
    /usr/local/bin/code-server --allow-http --no-auth --data-dir /data /code
  else
    exec "$@"
fi
