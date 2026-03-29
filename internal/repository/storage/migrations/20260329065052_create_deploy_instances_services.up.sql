CREATE TABLE deploy_instance_services (
    id SERIAL PRIMARY KEY,
    app_id INT REFERENCES deploy_instances(id) ON DELETE CASCADE,
    services_name TEXT NOT NULL,
    status TEXT DEFAULT 'Stopped',
    created_at TIMESTAMP DEFAULT NOW()
);