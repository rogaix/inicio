#!/bin/bash

BACKUP_DIR="/backups"

DATE=$(date +%F_%H-%M-%S)

if [[ -z "${MYSQL_HOST}" || -z "${MYSQL_USER}" || -z "${MYSQL_DATABASE}" ]]; then
    echo "error: one or more mysql variables are not defined."
    exit 1
fi

mkdir -p $BACKUP_DIR
docker exec -e MYSQL_PWD="${MYSQL_PASSWORD}" mysql mysqldump -h $MYSQL_HOST -u $MYSQL_USER $MYSQL_DATABASE > $BACKUP_DIR/mysql_backup_$DATE.sql

find $BACKUP_DIR -type f -mtime +7 -name '*.sql' -exec rm -- '{}' \;
echo "Backup finished at $DATE"
