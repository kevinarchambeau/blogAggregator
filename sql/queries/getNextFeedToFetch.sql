-- name: GetNextFeedToFetch :one
SELECT * from feeds
order by last_fetched_at nulls first
limit 1;