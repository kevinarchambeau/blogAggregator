-- +goose Up
CREATE TABLE feed_follows (
    id uuid primary key default gen_random_uuid(),
    created_at timestamp not null,
    updated_at timestamp not null,
    user_id uuid references users(id) on delete cascade not null,
    feed_id uuid references feeds(id) on delete cascade not null,
    unique (user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;
