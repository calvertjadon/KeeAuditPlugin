-- name: GetSpecificationByCode :one
select * from specifications
where code = $1;
