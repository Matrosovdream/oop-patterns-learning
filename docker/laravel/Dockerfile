# Laravel environment for the architecture + DDD/CQRS lessons.
FROM php:8.3-cli

# Extensions Laravel needs (plus sqlite/pgsql drivers for the examples).
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        git unzip libzip-dev libpq-dev libsqlite3-dev libonig-dev \
    && docker-php-ext-install zip pdo pdo_sqlite pdo_pgsql bcmath \
    && rm -rf /var/lib/apt/lists/*

COPY --from=composer:2 /usr/bin/composer /usr/bin/composer

WORKDIR /app

COPY docker/laravel/entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]

EXPOSE 8000
CMD ["php", "artisan", "serve", "--host=0.0.0.0", "--port=8000"]
