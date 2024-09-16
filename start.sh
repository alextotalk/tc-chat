#!/bin/bash

# Ініціалізуємо базу даних, якщо вона ще не існує
#if [ ! -d "$PGDATA" ]; then
#    mkdir -p "$PGDATA"
#    chown -R postgres:postgres "$PGDATA"
#    su - postgres -c "/usr/bin/initdb -D '$PGDATA'"
#fi
#
## Запускаємо PostgreSQL у фоновому режимі на порту 10000
##su - postgres -c "/usr/bin/pg_ctl -D '$PGDATA' -o '-p 10000' -w start"
#
## Чекаємо, поки PostgreSQL запуститься
#sleep 5

# Створюємо роль 'alex' з паролем 'secret'
#su - postgres -c "psql -p 10000 -c \"CREATE ROLE alex WITH LOGIN PASSWORD 'secret';\""

# Створюємо базу даних 'chat' з власником 'alex'
#su - postgres -c "psql -p 10000 -c \"CREATE DATABASE chat OWNER alex;\""

# Запускаємо Go-додаток
/app
