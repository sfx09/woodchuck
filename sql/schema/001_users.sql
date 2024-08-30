-- +goose Up
create table users (
  id uuid primary key,
  created_at timestamp,
  updated_at timestamp,
  name varchar(50) NOT NULL
);

-- +goose Down
DROP TABLE users;
