#!/bin/bash

# Ініціалізуємо базу даних, якщо вона ще не існує
if [ ! -d "$PGDATA" ]; then
    mkdir -p "$PGDATA"
    chown -R postgres:postgres "$PGDATA"
    su - postgres -c "/usr/bin/initdb -D '$PGDATA'"
fi

# Запускаємо PostgreSQL у фоновому режимі
su - postgres -c "/usr/bin/pg_ctl -D '$PGDATA' -w start"

# Чекаємо, поки PostgreSQL запуститься
sleep 5

# Виконуємо міграції або ініціалізацію бази даних, якщо потрібно
# Наприклад:
# su - postgres -c "psql -U $POSTGRES_USER -d $POSTGRES_DB -f /path/to/init.sql"

# Створюємо базу даних, якщо вона не існує
su - postgres -c "psql -c \"CREATE DATABASE chat;\""

# Запускаємо Go-додаток
/app
