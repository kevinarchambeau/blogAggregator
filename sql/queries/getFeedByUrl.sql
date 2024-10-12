-- name: GetFeedByUrl :one
SELECT * from feeds
    where url = $1;