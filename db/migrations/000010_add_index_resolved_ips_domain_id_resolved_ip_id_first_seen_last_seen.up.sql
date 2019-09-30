CREATE INDEX index_resolved_ips_on_domain_id
ON resolved_ips (domain_id);

CREATE INDEX index_resolved_ips_on_resolved_ip_id
ON resolved_ips (resolved_ip_id);

CREATE INDEX index_resolved_ips_on_first_seen
ON resolved_ips (first_seen);

CREATE INDEX index_resolved_ips_on_last_seen
ON resolved_ips (last_seen);