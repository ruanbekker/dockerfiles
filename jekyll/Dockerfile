FROM ruby:2.5-alpine

RUN apk add --no-cache build-base gcc bash cmake git
WORKDIR /site

RUN bundle init && \
    bundle install --path vendor/bundle && \
    bundle add jekyll jekyll-paginate kramdown pygments.rb && \
    bundle exec jekyll new --force --skip-bundle . && \
    bundle install

CMD [ "bundle", "exec", "jekyll", "serve", "--force_polling", "-H", "0.0.0.0", "-P", "4000" ]
