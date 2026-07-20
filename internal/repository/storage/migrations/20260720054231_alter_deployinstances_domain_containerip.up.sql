ALTER TABLE deploy_instances
ADD COLUMN domain_name VARCHAR(255),
ADD COLUMN container_ip VARCHAR(45);