-- +goose Up
-- +goose StatementBegin

CREATE TABLE rest_api_users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(50) UNIQUE,
    first_name VARCHAR(50),
    second_name VARCHAR(50),
    email VARCHAR(50) UNIQUE,
    password_hash VARCHAR(100)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS rest_api_users;
-- +goose StatementEnd
