FROM node:8.4.0-alpine

ARG GITBOOK_VERSION=3.2.3
LABEL version=$GITBOOK_VERSION

RUN npm install --global gitbook-cli \
        && gitbook fetch ${GITBOOK_VERSION} \
        && npm cache verify \
        && rm -rf /tmp/*

WORKDIR /srv/gitbook
VOLUME /srv/gitbook

EXPOSE 4000 35729

# https://github.com/GitbookIO/gitbook-cli/issues/55
COPY copyPluginAssets.js /root/.gitbook/versions/${GITBOOK_VERSION}/lib/output/website/copyPluginAssets.js

CMD ["gitbook", "--help"]
