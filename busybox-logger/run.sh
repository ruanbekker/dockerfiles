#!/usr/bin/env sh

while true
do
  sleep 1
  echo "[$(date +%FT%T)] level=info version=$version message=\"this is the $version container running with random id of $RANDOM\""
done
