ALTER TABLE resolved_domains
DROP CONSTRAINT fk_resolved_domain_id;

ALTER TABLE resolved_domains
DROP CONSTRAINT fk_domain_id;

ALTER TABLE resolved_domains
DROP CONSTRAINT fk_source_id;