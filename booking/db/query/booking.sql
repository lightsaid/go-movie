-- name: CreateBooking :one
insert into booking (
    "user_id",
    "show_id",
    "seat_number",
    "created_at",
    "status"
) values (
    $1, $2, $3, $4, $5
) returning *;

-- name: GetBooking :one
select * from booking where id = $1;

-- name: GetBookingByUserID :many
select * from booking where user_id = $1 limit $2 offset $3;

