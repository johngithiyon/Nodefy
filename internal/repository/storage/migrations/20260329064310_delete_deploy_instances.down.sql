-- up.sql
CREATE TABLE deploy_instances (
    id SERIAL PRIMARY KEY,
    appname TEXT,
    languages TEXT[],
    services TEXT[],
    username TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);