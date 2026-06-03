#!/bin/sh
# First-boot setup for the mounted Laravel app: install deps, create .env,
# generate the app key, and ensure the SQLite database file exists.
set -e
cd /app

if [ -f composer.json ] && [ ! -d vendor ]; then
  echo "[entrypoint] Installing Composer dependencies (first run, may take a minute)…"
  composer install --no-interaction --no-progress
fi

if [ -f .env.example ] && [ ! -f .env ]; then
  echo "[entrypoint] Creating .env from .env.example…"
  cp .env.example .env
fi

if [ -f artisan ] && [ -f .env ]; then
  grep -q '^APP_KEY=base64' .env 2>/dev/null || php artisan key:generate --force || true
fi

# SQLite store used by the DDD/CQRS persistence examples.
if [ ! -f database/database.sqlite ]; then
  mkdir -p database
  touch database/database.sqlite
fi

exec "$@"
