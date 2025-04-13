CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    passwordHash VARCHAR(100),
    password VARCHAR(100)
);

INSERT INTO users (id, username, email)
VALUES (1, 'firstUser', 'test@email.com');

INSERT INTO users (id, username, email)
VALUES (2, 'secondUser', 'test@email.com');