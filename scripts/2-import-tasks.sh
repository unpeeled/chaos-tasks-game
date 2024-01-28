#!/usr/bin/bash

pwd=$(pwd)
source ${pwd}/../etc/chaostasks-env.sh

if [ -z $1 ]; then
    echo "! Script needs an argument with path to text file for reading tasks."
    exit 1
fi

if [ -z $DB_HOST ] | [ -z $DB_NAME ] | [ -z $DB_USER ] | [ -z $DB_PASSWORD ]; then
    echo "! Setup script can only be used if DB_xx environment variables are set."
    exit 1
fi

while read line; do
    PGPASSWORD=$DB_PASSWORD psql -d $DB_NAME -h $DB_HOST -U $DB_USER -c "insert into tasks(task) VALUES('${line}')"
done < $1

