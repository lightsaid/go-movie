-- name: CreateCity :one
insert into city (
    name, lat, long
)values(
    $1, $2, $3
) returning *;

-- name: GetCities :many
select * from city limit $1 offset $2;

-- name: GetAllCity :many
select * from city;