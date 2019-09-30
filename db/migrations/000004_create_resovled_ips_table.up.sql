CREATE TABLE IF NOT EXISTS resolved_ips(
	id serial PRIMARY KEY,
	domain_id integer NOT NULL,
	resolved_ip_id integer NOT NULL,
	source_id integer NOT NULL,
	first_seen TIMESTAMP NOT NULL,
	last_seen TIMESTAMP NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);