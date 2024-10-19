// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: markFeedFetched.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const markFeedFetched = `-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = $1, updated_at = $1
    WHERE id = $2
`

type MarkFeedFetchedParams struct {
	LastFetchedAt sql.NullTime
	ID            uuid.UUID
}

func (q *Queries) MarkFeedFetched(ctx context.Context, arg MarkFeedFetchedParams) error {
	_, err := q.db.ExecContext(ctx, markFeedFetched, arg.LastFetchedAt, arg.ID)
	return err
}
