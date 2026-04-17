-- name: AttachAuditSpecs :exec
with input_codes as (
  select unnest($2::text[]) as code
)
insert into audit_specifications (audit_id,specification_id) 
select $1, s.id
from specifications s
join input_codes i on i.code = s.code;

