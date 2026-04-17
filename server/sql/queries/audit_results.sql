-- name: InsertAuditResult :one
insert into audit_results (
  id,
  audit_user_id,
  entry_name,
  entry_path
) values (
  gen_random_uuid(),
  $1,$2,$3
)
returning id;
