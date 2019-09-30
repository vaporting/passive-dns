CREATE TABLE IF NOT EXISTS ips(
	id serial PRIMARY KEY,
	ip bytea UNIQUE NOT NULL,
	type varchar (4) NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);