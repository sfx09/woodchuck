-- name: CreatePost :one 
INSERT INTO posts (id, created_at, updated_at, title, url, description, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetPostsByUser :many
SELECT * from posts
WHERE feed_id IN (
  SELECT feed_id from follows
  WHERE user_id = $1
); 
