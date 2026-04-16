-- name: GetAudit :one
select * from audits
where id = $1;

-- name: CreateAudit :one
insert into audits (
  id,
  created_at,
  updated_at
) values (
  $1,
  now(),
  now()
)
returning *;
