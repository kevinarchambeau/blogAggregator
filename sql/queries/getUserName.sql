-- name: GetUser :one
SELECT * from users
    where name = $1
;