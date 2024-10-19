-- +goose Up
create table posts (
    id uuid primary key default gen_random_uuid(),
    created_at timestamp not null,
    updated_at timestamp not null,
    title varchar not null,
    url varchar unique not null,
    description varchar not null,
    published_at timestamp not null,
    feed_id uuid references feeds(id) on delete cascade not null
);

-- +goose Down
DROP TABLE posts