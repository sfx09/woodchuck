-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, api_key)
VALUES ($1, $2, $3, $4, ENCODE(SHA256(RANDOM()::TEXT::BYTEA), 'HEX'))
RETURNING *;

-- name: GetUserByApiKey :one
SELECT * from users WHERE api_key = $1;
