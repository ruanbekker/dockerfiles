#!/bin/bash
set -e

if [ $# -eq 0 ]
  then
    /usr/local/bin/code-server --cert=/certs/localhost.crt --cert-key=/certs/localhost.key --user-data-dir /data /code
    #/usr/local/bin/code-server --cert=~/certs/localhost.crt --cert-key=~/certs/localhost.key --allow-http --no-auth --data-dir /data /code
  else
    exec "$@"
fi
