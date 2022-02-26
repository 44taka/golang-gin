#!/bin/sh
# wait-for-it.sh

set -e

until nc -z postgres 5432; do
  >&2 echo "postgres is unavailable - sleeping"
  sleep 3
done

>&2 echo "postgres is up - executing command"

# until nc -z mysql 3306; do
#   >&2 echo "mysql is unavailable - sleeping"
#   sleep 3
# done

# until nc -z mysql-test 33060; do
#   >&2 echo "mysql-test is unavailable - sleeping"
#   sleep 3
# done

# >&2 echo "mysql and mysql-test is up - executing command"

exec $@