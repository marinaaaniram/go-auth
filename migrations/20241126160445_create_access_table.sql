-- +goose Up
create table access_endpoint (
    id serial primary key,
    endpoint text not null,
    role text not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
drop table access_endpoint;
