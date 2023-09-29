-- +goose Up
ALTER TABLE users.users ADD CONSTRAINT phone_unique UNIQUE (phone);


-- +goose DOWN
ALTER TABLE users.users DROP CONSTRAINT phone_unique;
