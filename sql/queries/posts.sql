-- name: CreatePost :one 
INSERT INTO posts (id, created_at, updated_at, title, url, description, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
