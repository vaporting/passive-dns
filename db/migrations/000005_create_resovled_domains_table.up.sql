CREATE TABLE IF NOT EXISTS resolved_domains(
	id serial PRIMARY KEY,
	domain_id integer NOT NULL,
	resolved_domain_id integer NOT NULL,
	source_id integer NOT NULL,
	first_seen TIMESTAMP NOT NULL,
	last_seen TIMESTAMP NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);