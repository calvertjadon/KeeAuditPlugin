-- +goose Up
create table audits (
  id uuid primary key,
  created_at timestamp not null,
  updated_at timestamp not null
);

-- +goose Down
drop table audits;
