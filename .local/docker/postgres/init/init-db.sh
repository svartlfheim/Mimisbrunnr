#! /bin/bash

COMMANDS="\
CREATE ROLE ${MIGRATIONS_DB_USER} WITH NOINHERIT LOGIN PASSWORD '${MIGRATIONS_DB_USER_PASSWORD}';\
CREATE SCHEMA ${APP_SCHEMA} AUTHORIZATION ${MIGRATIONS_DB_USER};\
CREATE ROLE ${APP_DB_USER} WITH NOINHERIT LOGIN PASSWORD '${APP_DB_USER_PASSWORD}';\
GRANT USAGE ON SCHEMA ${APP_SCHEMA} to ${APP_DB_USER};
ALTER DEFAULT PRIVILEGES FOR USER ${MIGRATIONS_DB_USER} IN SCHEMA ${APP_SCHEMA} GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO ${APP_DB_USER};
"

echo "$COMMANDS" | PGPASSWORD=$POSTGRES_PASSWORD psql -U $POSTGRES_USER
