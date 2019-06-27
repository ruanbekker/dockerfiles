#!/usr/bin/env sh
HUGO_HOME=${HUGO_HOME:-/usr/local/hugo}
HUGO_THEME=${HUGO_THEME:-introduction}
PROJECT_PATH=${PROJECT_PATH:-/app}
PROJECT_URL=${PROJECT_URL:-https://0.0.0.0}

cd /app
rm -rf /app/*
git clone ${GITHUB_URL} .

hugo server \
  --themesDir=${PROJECT_PATH}/themes \
  --theme=${HUGO_THEME} \
  --bind=0.0.0.0 \
  --baseUrl=${PROJECT_URL}/ \
  --port=8080 \
  --appendPort=false \
  --buildDrafts \
  --environment production
