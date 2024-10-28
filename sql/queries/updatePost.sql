-- name: UpdatePost :exec
UPDATE posts SET
 updated_at = $1, title = $2, description = $3
WHERE url = $4;