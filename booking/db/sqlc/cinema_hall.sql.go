// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: cinema_hall.sql

package db

import (
	"context"
)

const createCinemaHall = `-- name: CreateCinemaHall :one
insert into cinema_hall (
    "cinema_id",
    "name",
    "total_seats"
) values (
  $1, $2, $3
) returning id, cinema_id, name, total_seats
`

type CreateCinemaHallParams struct {
	CinemaID   int64  `json:"cinema_id"`
	Name       string `json:"name"`
	TotalSeats int32  `json:"total_seats"`
}

func (q *Queries) CreateCinemaHall(ctx context.Context, arg CreateCinemaHallParams) (CinemaHall, error) {
	row := q.db.QueryRowContext(ctx, createCinemaHall, arg.CinemaID, arg.Name, arg.TotalSeats)
	var i CinemaHall
	err := row.Scan(
		&i.ID,
		&i.CinemaID,
		&i.Name,
		&i.TotalSeats,
	)
	return i, err
}

const getCinemaHallList = `-- name: GetCinemaHallList :many
select id, cinema_id, name, total_seats from cinema_hall limit $1 offset $2
`

type GetCinemaHallListParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetCinemaHallList(ctx context.Context, arg GetCinemaHallListParams) ([]CinemaHall, error) {
	rows, err := q.db.QueryContext(ctx, getCinemaHallList, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CinemaHall{}
	for rows.Next() {
		var i CinemaHall
		if err := rows.Scan(
			&i.ID,
			&i.CinemaID,
			&i.Name,
			&i.TotalSeats,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCinemaHall = `-- name: UpdateCinemaHall :one
update cinema_hall set
  "cinema_id" = $2,
  "name" = $3,
  "total_seats" = $4
where id = $1 returning id, cinema_id, name, total_seats
`

type UpdateCinemaHallParams struct {
	ID         int64  `json:"id"`
	CinemaID   int64  `json:"cinema_id"`
	Name       string `json:"name"`
	TotalSeats int32  `json:"total_seats"`
}

func (q *Queries) UpdateCinemaHall(ctx context.Context, arg UpdateCinemaHallParams) (CinemaHall, error) {
	row := q.db.QueryRowContext(ctx, updateCinemaHall,
		arg.ID,
		arg.CinemaID,
		arg.Name,
		arg.TotalSeats,
	)
	var i CinemaHall
	err := row.Scan(
		&i.ID,
		&i.CinemaID,
		&i.Name,
		&i.TotalSeats,
	)
	return i, err
}
