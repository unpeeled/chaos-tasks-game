#!/usr/bin/bash

if [ -t 0 ]; then
    echo "! Script needs STDIN with tasks to be imported."
    exit 1
fi

if [ -z $DB_HOST ] | [ -z $DB_NAME ] | [ -z $DB_USER ]; then
    echo "! Setup script can only be used if DB_xx environment variables are set."
    exit 1
fi

if [ -n $DB_PASSWORD_FILE ];  then
    export PGPASSWORD=$(cat $DB_PASSWORD_FILE)
elif [-n $DB_PASSWORD ]; then
    export PGPASSWORD=$DB_PASSWORD
else
    echo "No Password is set"
fi

while read line; do
    /usr/bin/psql -d $DB_NAME -h $DB_HOST -U $DB_USER -c "insert into tasks(task) VALUES('${line}')"
done < "${1:-/dev/stdin}"

