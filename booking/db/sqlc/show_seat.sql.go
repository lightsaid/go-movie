// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: show_seat.sql

package db

import (
	"context"
)

const createShowSeat = `-- name: CreateShowSeat :one
insert into show_seat (
    "cinema_seat_id",
    "show_id",
    "booking_id",
    "status",
    "price"
) values (
    $1, $2, $3, $4, $5
) returning id, cinema_seat_id, show_id, booking_id, status, price
`

type CreateShowSeatParams struct {
	CinemaSeatID int64  `json:"cinema_seat_id"`
	ShowID       int64  `json:"show_id"`
	BookingID    int64  `json:"booking_id"`
	Status       int32  `json:"status"`
	Price        string `json:"price"`
}

func (q *Queries) CreateShowSeat(ctx context.Context, arg CreateShowSeatParams) (ShowSeat, error) {
	row := q.db.QueryRowContext(ctx, createShowSeat,
		arg.CinemaSeatID,
		arg.ShowID,
		arg.BookingID,
		arg.Status,
		arg.Price,
	)
	var i ShowSeat
	err := row.Scan(
		&i.ID,
		&i.CinemaSeatID,
		&i.ShowID,
		&i.BookingID,
		&i.Status,
		&i.Price,
	)
	return i, err
}

const getShowSeat = `-- name: GetShowSeat :one
select id, cinema_seat_id, show_id, booking_id, status, price from show_seat where id = $1
`

func (q *Queries) GetShowSeat(ctx context.Context, id int64) (ShowSeat, error) {
	row := q.db.QueryRowContext(ctx, getShowSeat, id)
	var i ShowSeat
	err := row.Scan(
		&i.ID,
		&i.CinemaSeatID,
		&i.ShowID,
		&i.BookingID,
		&i.Status,
		&i.Price,
	)
	return i, err
}
