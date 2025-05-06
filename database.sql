CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    passwordHash VARCHAR(100),
    password VARCHAR(100)
);

CREATE TABLE tg_users
(
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    tg_id int UNIQUE NOT NULL,
    role_id int
);

CREATE TABLE roles
(
    id SERIAL PRIMARY KEY,
    role_desc VARCHAR(100)
);