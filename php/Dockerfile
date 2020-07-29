FROM php:7.4-cli

WORKDIR /app

ADD . /app

RUN apt-get update && apt-get install -y git unzip

RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer && \
    composer install

RUN php ./vendor/bin/phpunit /app/test

CMD php -S 0.0.0.0:8080
