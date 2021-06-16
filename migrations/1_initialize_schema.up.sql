CREATE TABLE IF NOT EXISTS prr_users (
	id serial PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	password VARCHAR ( 255 ) NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
	deleted_at TIMESTAMP DEFAULT NULL, 	
    last_login TIMESTAMP 
);

CREATE TABLE IF NOT EXISTS prr_groups (
	id serial PRIMARY KEY,
	name VARCHAR ( 50 ) UNIQUE NOT NULL,
    protected BOOL NOT NULL DEFAULT false,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
	deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS prr_users_has_groups (
	id serial PRIMARY KEY,
	id_user serial NOT NULL,
	id_group serial NOT NULL,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
	deleted_at TIMESTAMP DEFAULT NULL, 
	FOREIGN KEY (id_user) 	REFERENCES prr_users (id),
	FOREIGN KEY (id_group) 	REFERENCES prr_groups (id)
);

CREATE TABLE IF NOT EXISTS prr_products (
	id VARCHAR  PRIMARY KEY,
	name VARCHAR ( 50 ) UNIQUE NOT NULL,
    description TEXT DEFAULT NULL,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
	deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS prr_teams (
	id serial PRIMARY KEY,
	name VARCHAR ( 50 ) UNIQUE NOT NULL,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
	deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS prr_users_has_teams (
	id serial PRIMARY KEY,
	id_user serial NOT NULL,
	id_team serial NOT NULL,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
	deleted_at TIMESTAMP DEFAULT NULL, 	
    last_login TIMESTAMP,
	FOREIGN KEY (id_user) 	REFERENCES prr_users (id),
	FOREIGN KEY (id_team) 	REFERENCES prr_teams (id)
);

CREATE TABLE IF NOT EXISTS prr_apps (
	id serial PRIMARY KEY,
	name VARCHAR ( 50 ) UNIQUE NOT NULL,
	id_team serial NOT NULL, 
	id_product VARCHAR NOT NULL, 
	language VARCHAR ( 50 ), 
	framework VARCHAR ( 50 ), 
	architecture VARCHAR ( 50 ), 
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
	deleted_at TIMESTAMP DEFAULT NULL, 
	FOREIGN KEY (id_team) 		REFERENCES prr_teams (id),
	FOREIGN KEY (id_product)	REFERENCES prr_products (id)
);