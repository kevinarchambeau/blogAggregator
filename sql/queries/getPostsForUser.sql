-- name: GetPostsForUsers :many
select * from posts where
    feed_id in (
        select id from feeds where user_id = $1
    )
    order by published_at desc
    limit $2;
