#!/bin/sh
set -e

host="$DB_HOST"
port="$DB_PORT"

echo "⏳ Waiting for PostgreSQL at $host:$port..."

until nc -z "$host" "$port"; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

echo "✅ PostgreSQL is up - starting application"

exec "$@"
