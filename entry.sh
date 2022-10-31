#!/bin/sh

until nc -z -v -w30 productdb 5432
do
  echo "Waiting for database connection..."
  # wait for 5 seconds before check again
  sleep 5
done

echo -e "\e[34m >>> Starting the server \e[97m"
$1
