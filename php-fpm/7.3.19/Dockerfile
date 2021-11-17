FROM php:7.3.19-fpm-alpine AS target

WORKDIR /var/www/html

RUN apk --no-cache add \
    build-base \
    freetype-dev \
    libjpeg-turbo-dev \
    libpng-dev \
    libzip-dev \
    libxml2-dev \
    zip \
    jpegoptim optipng pngquant gifsicle \
    vim \
    unzip \
    git \
    curl \
    supervisor

RUN docker-php-ext-install pdo_mysql mbstring zip exif pcntl bcmath json ctype fileinfo tokenizer
RUN docker-php-ext-configure gd --with-gd --with-freetype-dir=/usr/include/ --with-jpeg-dir=/usr/include/ --with-png-dir=/usr/include/
RUN docker-php-ext-install gd

RUN apk add autoconf && pecl install -o -f redis \
    && rm -rf /tmp/pear \
    && docker-php-ext-enable redis && apk del autoconf
