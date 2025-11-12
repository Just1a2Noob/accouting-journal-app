-- +goose Up
CREATE TABLE users(
    id uuid primary key,
    username TEXT NOT NULL,
    hash_pass TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    email TEXT NOT NULL UNIQUE,
    section TEXT NOT NULL DEFAULT 'unset',
    role TEXT NOT NULL DEFAULT 'employee'
);

-- +goose Down
DROP TABLE users;
