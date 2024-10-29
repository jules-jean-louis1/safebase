const { Client } = require('pg');

const pgclient = new Client({
    host: process.env.POSTGRES_HOST,
    port: process.env.POSTGRES_PORT,
    user: process.env.POSTGRES_USER,
    password: process.env.POSTGRES_PASSWORD,
    database: process.env.POSTGRES_DB,
});

pgclient.connect();

const createTables = `

CREATE TYPE IF NOT EXISTS user_role AS ENUM ('admin', 'user');
CREATE TYPE IF NOT EXISTS backup_status AS ENUM ('pending', 'in_progress', 'success', 'failed');
CREATE TYPE IF NOT EXISTS backup_type AS ENUM ('manual', 'scheduled');
CREATE TYPE IF NOT EXISTS restore_status AS ENUM ('pending', 'in_progress', 'success', 'failed');

CREATE TABLE IF NOT EXISTS database (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    type VARCHAR NOT NULL,
    host VARCHAR NOT NULL,
    port VARCHAR NOT NULL,
    username VARCHAR NOT NULL,
    password VARCHAR,
    database_name VARCHAR NOT NULL,
    is_cron_active BOOLEAN,
    cron_schedule VARCHAR,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS backup (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    database_id UUID REFERENCES database (id) ON DELETE SET NULL,
    status backup_status NOT NULL,
    backup_type backup_type NOT NULL,
    filename VARCHAR NOT NULL,
    size VARCHAR,
    error_msg TEXT,
    log TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS restore (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    database_id UUID REFERENCES database (id) ON DELETE SET NULL,
    backup_id UUID REFERENCES backup (id) ON DELETE SET NULL,
    status restore_status NOT NULL,
    filename VARCHAR NOT NULL,
    error_msg TEXT,
    log TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

`;

pgclient.query(createTables, (err, res) => {
    if (err) {
        console.error('Error creating tables:', err);
    } else {
        console.log('Tables created successfully');
    }
    pgclient.end();
});
