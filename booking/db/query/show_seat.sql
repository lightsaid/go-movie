-- name: CreateShowSeat :one
insert into show_seat (
    "cinema_seat_id",
    "show_id",
    "booking_id",
    "status",
    "price"
) values (
    $1, $2, $3, $4, $5
) returning *;

-- name: GetShowSeat :one
select * from show_seat where id = $1;

