-- name: CreateCinema :one 
insert into cinema (
  "city_id",
  "name",
  "lat",
  "long",
  "total_cinema_halls"
)values (
    $1, $2, $3, $4, $5
) returning *;

-- name: GetCinemaList :many
select * from cinema limit $1 offset $2;