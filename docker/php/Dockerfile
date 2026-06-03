# Plain PHP 8.3 CLI environment for the OOP / SOLID / GoF / pure-PHP DDD lessons.
FROM php:8.3-cli

# Composer (for the PSR-4 autoloader; the examples have no heavy dependencies).
COPY --from=composer:2 /usr/bin/composer /usr/bin/composer

# Small quality-of-life extension used by some examples.
RUN apt-get update \
    && apt-get install -y --no-install-recommends git unzip \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY docker/php/entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]

CMD ["php", "-a"]
