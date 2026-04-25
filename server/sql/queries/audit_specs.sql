-- name: AttachAuditSpecs :exec
insert into audit_specifications (audit_id,specification_id, threshold) 
select $1, $2, $3;

