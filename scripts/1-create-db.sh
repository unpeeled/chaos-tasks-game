#!/usr/bin/bash

pwd=$(pwd)
source ${pwd}/../etc/chaostasks-env.sh

if [ -z $DB_HOST ] | [ -z $DB_NAME ] | [ -z $DB_USER ] | [ -z $DB_PASSWORD ]; then
    echo "! Setup script can only be used if DB_xx variables are set."
    exit 1
fi

sudo -i -u postgres createuser $DB_USER
sudo -i -u postgres createdb $DB_NAME
sudo -i -u postgres psql -c "alter user ${DB_USER} with encrypted password '${DB_PASSWORD}'"
sudo -i -u postgres psql -c "alter database ${DB_NAME} OWNER TO ${DB_USER}"

PGPASSWORD=$DB_PASSWORD psql -d $DB_NAME -h $DB_HOST -U $DB_USER -f ${pwd}/create-pg-db.sql

