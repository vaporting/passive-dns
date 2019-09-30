ALTER TABLE resolved_ips
ADD CONSTRAINT fk_resolved_ip_id
FOREIGN KEY (resolved_ip_id)
REFERENCES ips (id);

ALTER TABLE resolved_ips
ADD CONSTRAINT fk_domain_id
FOREIGN KEY (domain_id)
REFERENCES domains (id);

ALTER TABLE resolved_ips
ADD CONSTRAINT fk_source_id
FOREIGN KEY (source_id)
REFERENCES sources (id);