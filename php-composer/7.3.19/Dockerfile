FROM php:7.3.19-cli

RUN apt-get update \
    && apt-get install \
       default-mysql-client \
       default-libmysqlclient-dev \
       curl git libfreetype6-dev \
       libjpeg62-turbo-dev \
       libmcrypt-dev libpng-dev \
       libzip-dev libxml2 libxml2-dev libcurl4-gnutls-dev libssl-dev sqlite3 libsqlite3-dev libedit-dev libxslt1-dev -y

RUN docker-php-ext-install iconv \
    && pecl install mcrypt-1.0.3 \
    && docker-php-ext-enable mcrypt \
    && docker-php-ext-configure gd --with-freetype-dir=/usr/include/ --with-jpeg-dir=/usr/include/ \
    && docker-php-ext-install gd \
    && docker-php-ext-install zip \
    && docker-php-ext-install mysqli \
    && docker-php-ext-install opcache \
    && docker-php-ext-install mbstring \
    && docker-php-ext-install bcmath \
    && docker-php-ext-install pcntl \
    && docker-php-ext-install calendar \
    && docker-php-ext-install pdo \
    && docker-php-ext-install curl \
    && docker-php-ext-install exif \
    && docker-php-ext-install gettext \
    && docker-php-ext-install gmp \
    && docker-php-ext-install exif \
    && docker-php-ext-install intl \
    && docker-php-ext-install pdo_mysql \
    && docker-php-ext-install pdo_sqlite \
    && docker-php-ext-install readline \
    && pecl install -o -f redis \
    && rm -rf /tmp/pear && docker-php-ext-enable redis \
    && docker-php-ext-install shmop \
    && docker-php-ext-install sockets \
    && docker-php-ext-install sysvsem \
    && docker-php-ext-install sysvshm \
    && docker-php-ext-install wddx \
    && docker-php-ext-install xsl \
    && pecl install msgpack && docker-php-ext-enable msgpack \
    && pecl install igbinary && docker-php-ext-enable igbinary \
    && docker-php-ext-install sysvmsg

#RUN curl -sS https://getcomposer.org/installer | php && mv composer.phar /usr/local/bin/composer
RUN curl -O "https://getcomposer.org/download/1.10.21/composer.phar" \
    && chmod a+x composer.phar \
    && mv composer.phar /usr/bin/composer
RUN composer global require phpunit/phpunit "^7.5"
RUN composer global require mybuilder/phpunit-accelerator "^1.2" --dev

ENV PATH /root/.composer/vendor/bin:$PATH
RUN ln -s /root/.composer/vendor/bin/phpunit /usr/bin/phpunit
