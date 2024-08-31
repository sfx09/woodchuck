-- +goose Up
create table users (
  id uuid primary key,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  name varchar(50) NOT NULL
);

-- +goose Down
DROP TABLE users;
