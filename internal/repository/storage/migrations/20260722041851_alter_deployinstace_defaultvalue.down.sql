ALTER TABLE deploy_instances
ALTER COLUMN domain_name DROP DEFAULT;

ALTER TABLE deploy_instances
ALTER COLUMN port_number DROP DEFAULT;