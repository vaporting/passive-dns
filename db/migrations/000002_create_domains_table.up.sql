CREATE TABLE IF NOT EXISTS domains(
	id serial PRIMARY KEY,
	name text UNIQUE NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);
