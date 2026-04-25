-- +goose Up
insert into specifications (id, code, description) values (
  gen_random_uuid(),
  'entropy.min',
  'specifies the minimum allowed entropy'
);
insert into specifications (id, code, description) values (
  gen_random_uuid(),
  'duplicates.max',
  'specifies that a given password may only appear a specified number of times in a database'
);

-- +goose Down
delete from specifications
where code = 'entropy.min' or code = 'duplicates.max';
