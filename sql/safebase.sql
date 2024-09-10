-- Création du type ENUM pour les rôles d'utilisateur
CREATE TYPE user_role AS ENUM ('admin', 'user');

-- Création du type ENUM pour le statut des backups
CREATE TYPE backup_status AS ENUM ('pending', 'in_progress', 'success', 'failed');

-- Création du type ENUM pour le type de backup
CREATE TYPE backup_type AS ENUM ('manual', 'scheduled');

-- Création du type ENUM pour le statut des restaurations
CREATE TYPE restore_status AS ENUM ('pending', 'in_progress', 'success', 'failed');

-- Création de la table des utilisateurs
CREATE TABLE "user" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    role user_role NOT NULL DEFAULT 'user',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Création de la table des bases de données
CREATE TABLE "database" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    type VARCHAR NOT NULL,
    host VARCHAR NOT NULL,
    port VARCHAR NOT NULL,
    username VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    database_name VARCHAR NOT NULL,
    connection_string VARCHAR,
    cron_schedule VARCHAR,  -- Ajout du planning de backup ici
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Création de la table des backups
CREATE TABLE "backup" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    database_id UUID NOT NULL REFERENCES database (id) ON DELETE CASCADE,
    status backup_status NOT NULL,
    backup_type backup_type NOT NULL,
    filename VARCHAR NOT NULL,
    size VARCHAR,
    error_msg TEXT,
    log TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Création de la table des restaurations
CREATE TABLE "restore" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    database_id UUID NOT NULL REFERENCES database (id) ON DELETE CASCADE,
    backup_id UUID NOT NULL REFERENCES backup (id) ON DELETE SET NULL,
    status restore_status NOT NULL,
    filename VARCHAR NOT NULL,
    error_msg TEXT,
    log TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);