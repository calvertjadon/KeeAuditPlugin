-- +goose Up
create table audit_specifications (
  audit_id uuid not null references audits(id) on delete cascade,
  specifications_id uuid not null references specifications(id),

  primary key (audit_id, specifications_id)
);

-- +goose Down
drop table audit_specifications;
