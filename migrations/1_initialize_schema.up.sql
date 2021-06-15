CREATE TABLE IF NOT EXISTS prr_users (
	id serial PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	password VARCHAR ( 255 ) NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP 
);

CREATE TABLE IF NOT EXISTS prr_groups (
	id serial PRIMARY KEY,
	label VARCHAR ( 50 ) UNIQUE NOT NULL,
	created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP 
);

CREATE TABLE IF NOT EXISTS prr_users_has_groups (
	id serial PRIMARY KEY,
	id_user serial NOT NULL,
	id_group serial NOT NULL,
	created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP,
	FOREIGN KEY (id_user) 	REFERENCES prr_users (id),
	FOREIGN KEY (id_group) 	REFERENCES prr_groups (id)
);

CREATE TABLE IF NOT EXISTS prr_product (
	id serial PRIMARY KEY,
	name VARCHAR ( 50 ) UNIQUE NOT NULL,
	created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP 
);

CREATE TABLE IF NOT EXISTS prr_team (
	id serial PRIMARY KEY,
	name VARCHAR ( 50 ) UNIQUE NOT NULL,
	created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP 
);

CREATE TABLE IF NOT EXISTS prr_users_has_teams (
	id serial PRIMARY KEY,
	id_user serial NOT NULL,
	id_team serial NOT NULL,
	created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP,
	FOREIGN KEY (id_user) 	REFERENCES prr_users (id),
	FOREIGN KEY (id_team) 	REFERENCES prr_team (id)
);

CREATE TABLE IF NOT EXISTS prr_app (
	id serial PRIMARY KEY,
	name VARCHAR ( 50 ) UNIQUE NOT NULL,
	id_team serial NOT NULL, 
	id_product serial NOT NULL, 
	language VARCHAR ( 50 ), 
	framework VARCHAR ( 50 ), 
	architecture VARCHAR ( 50 ), 
	created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP, 
	FOREIGN KEY (id_team) 		REFERENCES prr_team (id),
	FOREIGN KEY (id_product)	REFERENCES prr_product (id)
);