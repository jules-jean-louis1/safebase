#!/bin/bash
set -e

# Créer plusieurs bases de données
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE music_library;
    CREATE DATABASE dev_library;
    CREATE DATABASE prod_library;
EOSQL

# Ensuite, exécuter les scripts SQL spécifiques à chaque base
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "music_library" -f /docker-entrypoint-initdb.d/init_db.sql
# Tu peux ajouter d'autres commandes similaires pour d'autres bases de données si nécessaire
