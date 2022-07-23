-- name: CreateCinemaHall :one
insert into cinema_hall (
    "cinema_id",
    "name",
    "total_seats"
) values (
  $1, $2, $3
) returning *;

-- name: GetCinemaHallList :many
select * from cinema_hall limit $1 offset $2;