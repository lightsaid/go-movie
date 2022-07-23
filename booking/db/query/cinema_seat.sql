-- name: CreateCinemaSeat :one
insert into cinema_seat (
    "cinema_hall_id", 
    "type",
    "seat_number"
) values (
    $1, $2, $3
) returning *;

-- name: GetCinemaSeatByHallID :many
select * from cinema_seat where cinema_hall_id = $1;

-- name: UpdateCinemaSeat :one 
update cinema_seat set
    type = $2,
    seat_number = $3
where id = $1 returning *; 
