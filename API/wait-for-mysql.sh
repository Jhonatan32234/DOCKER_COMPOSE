#!/bin/sh

set -e

host="$1"
shift

while ! mysql -h "$host" -u root -p"rootpass" -e "SELECT 1"; do
  echo "Esperando a que MySQL esté listo en $host..."
  sleep 2
done


exec "$@"
