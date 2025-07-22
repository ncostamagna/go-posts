-- name: UpdatePost :exec
UPDATE posts
SET title = $2, content = $3
WHERE id = $1;