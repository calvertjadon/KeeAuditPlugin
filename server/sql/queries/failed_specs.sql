-- name: InsertFailedSpec :exec
insert into audit_result_failed_specs (
  audit_result_id,
  specification_id
) values ($1, $2);
