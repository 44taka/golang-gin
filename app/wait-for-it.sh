#!/bin/sh
# wait-for-it.sh

set -e

until nc -z postgres 5432; do
  >&2 echo "postgres is unavailable - sleeping"
  sleep 3
done

until nc -z postgres-test 5432; do
  >&2 echo "postgres-test is unavailable - sleeping"
  sleep 3
done

>&2 echo "postgres and postgres-test is up - executing command"

exec $@