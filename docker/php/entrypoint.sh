#!/bin/sh
# Generate the Composer autoloader on first run (vendor/ is git-ignored).
set -e

if [ -f /app/composer.json ] && [ ! -d /app/vendor ]; then
  echo "[entrypoint] Installing Composer autoloader…"
  composer install --no-interaction --no-progress --quiet || \
    composer dump-autoload --quiet || true
fi

exec "$@"
