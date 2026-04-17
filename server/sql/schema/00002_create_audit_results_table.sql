-- +goose Up
create table users (
  id uuid primary key,
  username text unique not null,
  created_at timestamp not null,
  updated_at timestamp not null
);

create table specifications (
  id uuid primary key,
  code text unique not null,
  description text not null
);

create table audit_users (
  id uuid primary key,
  audit_id  uuid not null references audits(id) on delete cascade,
  user_id uuid not null references users(id) on delete cascade,
  completed_at timestamp not null,
  unique (audit_id, user_id)
);

create table audit_results (
  id uuid primary key,
  audit_user_id uuid not null references audit_users(id) on delete cascade,
  entry_id uuid not null,
  entry_name text not null,
  entry_path text not null
);

create table audit_result_failed_specs (
  audit_result_id uuid not null references audit_results(id) on delete cascade,
  specifications_id uuid not null references specifications(id),

  primary key (audit_result_id, specifications_id)
);

-- +goose Down
drop table audit_result_failed_specs;
drop table audit_results;
drop table audit_users;
drop table specifications;
drop table users;
