-- name: GetAudit :one
select * from audits
where id = $1;

-- name: CreateAudit :one
insert into audits (
  id,
  created_at,
  updated_at
) values (
  gen_random_uuid(),
  now(),
  now()
)
returning id;

