CREATE INDEX index_domains_on_name
ON domains (name);

CREATE INDEX gin_index_domains_on_name
ON domains USING gin (to_tsvector('english', name));