-- name: getUser :one
SELECT * from users
    where name = $1
;