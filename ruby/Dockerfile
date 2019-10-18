FROM alpine:3.9

RUN apk add --no-cache bash git wget curl vim build-base readline-dev openssl-dev zlib-dev
ENV PATH /usr/local/rbenv/shims:/usr/local/rbenv/bin:$PATH
ENV RBENV_ROOT /usr/local/rbenv
ENV RUBY_VERSION 2.5.1
ENV CONFIGURE_OPTS --disable-install-doc

RUN apk add --update linux-headers imagemagick-dev libffi-dev libffi-dev && rm -rf /var/cache/apk/* \
    && git clone --depth 1 git://github.com/sstephenson/rbenv.git ${RBENV_ROOT} \
    && git clone --depth 1 https://github.com/sstephenson/ruby-build.git ${RBENV_ROOT}/plugins/ruby-build \
    && git clone --depth 1 git://github.com/jf/rbenv-gemset.git ${RBENV_ROOT}/plugins/rbenv-gemset \
    && ${RBENV_ROOT}/plugins/ruby-build/install.sh \
    && echo 'eval "$(rbenv init -)"' >> /etc/profile.d/rbenv.sh \
    && rbenv install $RUBY_VERSION \
    && rbenv global $RUBY_VERSION \
    && rbenv rehash \
    && gem install bundler

COPY docker-entrypoint.sh /
RUN chmod +x /docker-entrypoint.sh

ENTRYPOINT ["/docker-entrypoint.sh"]
