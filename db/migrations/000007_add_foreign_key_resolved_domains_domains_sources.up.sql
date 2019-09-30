ALTER TABLE resolved_domains
ADD CONSTRAINT fk_resolved_domain_id
FOREIGN KEY (resolved_domain_id)
REFERENCES domains (id);

ALTER TABLE resolved_domains
ADD CONSTRAINT fk_domain_id
FOREIGN KEY (domain_id)
REFERENCES domains (id);

ALTER TABLE resolved_domains
ADD CONSTRAINT fk_source_id
FOREIGN KEY (source_id)
REFERENCES sources (id);