CREATE TABLE deploy_instances (
    id SERIAL PRIMARY KEY,
    appname TEXT NOT NULL UNIQUE,
    languages TEXT[],
    services TEXT[],
    username TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);