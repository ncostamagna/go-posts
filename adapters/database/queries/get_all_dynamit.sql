-- name: ListPosts :many
SELECT *
FROM posts
WHERE 
  ($1 IS NULL OR title = $1)
  AND
  ($2 IS NULL OR content = $2);