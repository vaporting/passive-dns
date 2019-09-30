CREATE INDEX index_resolved_domains_on_domain_id
ON resolved_domains (domain_id);

CREATE INDEX index_resolved_domains_on_resolved_domain_id
ON resolved_domains (resolved_domain_id);

CREATE INDEX index_resolved_domains_on_first_seen
ON resolved_domains (first_seen);

CREATE INDEX index_resolved_domains_on_last_seen
ON resolved_domains (last_seen);