-- name: GetUserName :one
SELECT name from users
    where id = $1
;