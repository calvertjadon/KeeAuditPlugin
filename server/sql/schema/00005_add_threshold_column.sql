-- +goose Up
alter table audit_specifications add column threshold int default 0 not null;

-- +goose Down
alter table audit_specifications drop column threshold;
