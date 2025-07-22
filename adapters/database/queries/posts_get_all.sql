-- name: GetAllPosts :many
SELECT id, title, content, created_at, updated_at FROM posts
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;
