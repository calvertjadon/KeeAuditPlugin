-- name: GetSpecificationByCode :one
select * from specifications
where code = $1;

-- name: GetSpecificationsByCodes :many
select * from specifications
where code = ANY($1::text[]);

-- name: CreateSpecification :one
insert into specifications (
  id,
  code,
  description
) values (
  gen_random_uuid(),
  $1,
  $2
)
returning id;
