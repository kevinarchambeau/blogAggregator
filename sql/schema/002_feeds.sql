-- +goose Up
CREATE TABLE feeds (
    id uuid primary key default gen_random_uuid(),
    created_at timestamp not null,
    updated_at timestamp not null,
    name varchar not null,
    url varchar unique not null,
    user_id uuid references users(id) on delete cascade
);

-- +goose Down
DROP TABLE feeds;
