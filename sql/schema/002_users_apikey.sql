-- +goose Up
alter table users add column api_key varchar(64) unique not null;
-- +goose Down
alter table users drop column api_key;
