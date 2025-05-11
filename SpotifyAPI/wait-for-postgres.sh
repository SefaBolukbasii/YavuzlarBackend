#!/bin/sh

set -e

host="$1"
shift
cmd="$@"

echo "Checking if PostgreSQL at $host:5432 is available..."

until pg_isready -h "$host" -p 5432 > /dev/null 2>&1; do
  echo "PostgreSQL is unavailable - waiting..."
  sleep 1
done

echo "PostgreSQL is up - executing command: $cmd"
exec $cmd
