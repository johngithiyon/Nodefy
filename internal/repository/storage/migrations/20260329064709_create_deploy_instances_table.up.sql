-- up.sql
CREATE TABLE deploy_instances (
    id SERIAL PRIMARY KEY,
    appname TEXT UNIQUE NOT NULL,
    username TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);