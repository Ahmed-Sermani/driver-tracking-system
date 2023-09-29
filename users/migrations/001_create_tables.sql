-- +goose Up
CREATE SCHEMA users;

SET
SEARCH_PATH TO users, PUBLIC;

CREATE TABLE users (
  id         text        NOT NULL,
  name       text        NOT NULL,
  phone      text        NOT NULL,
  enabled    bool        NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id)
);

-- +goose Down
DROP SCHEMA IF EXISTS users CASCADE;
-- DROP TABLE IF EXISTS users;
