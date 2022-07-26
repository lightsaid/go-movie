-- name: CreatePayment :one
insert into payment (
    booking_id, amount
) values ($1, $2) returning *;

-- name: GetPayment :one
select * from payment where id = $1;

-- name: GetPaymentByBookingID :one
select * from payment where booking_id = $1;