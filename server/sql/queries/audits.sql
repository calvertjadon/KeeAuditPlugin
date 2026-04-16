-- name: GetAudit :one
select * from audits
where id = $1;
