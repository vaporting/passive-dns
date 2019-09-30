ALTER TABLE resolved_ips
DROP CONSTRAINT fk_resolved_ip_id;

ALTER TABLE resolved_ips
DROP CONSTRAINT fk_domain_id;

ALTER TABLE resolved_ips
DROP CONSTRAINT fk_source_id;