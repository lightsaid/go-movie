// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: city.sql

package db

import (
	"context"
)

const createCity = `-- name: CreateCity :one
insert into city (
    name, lat, long
)values(
    $1, $2, $3
) returning id, name, lat, long
`

type CreateCityParams struct {
	Name string `json:"name"`
	Lat  string `json:"lat"`
	Long string `json:"long"`
}

func (q *Queries) CreateCity(ctx context.Context, arg CreateCityParams) (City, error) {
	row := q.db.QueryRowContext(ctx, createCity, arg.Name, arg.Lat, arg.Long)
	var i City
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Lat,
		&i.Long,
	)
	return i, err
}

const getAllCity = `-- name: GetAllCity :many
select id, name, lat, long from city
`

func (q *Queries) GetAllCity(ctx context.Context) ([]City, error) {
	rows, err := q.db.QueryContext(ctx, getAllCity)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []City{}
	for rows.Next() {
		var i City
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Lat,
			&i.Long,
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

const getCities = `-- name: GetCities :many
select id, name, lat, long from city limit $1 offset $2
`

type GetCitiesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetCities(ctx context.Context, arg GetCitiesParams) ([]City, error) {
	rows, err := q.db.QueryContext(ctx, getCities, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []City{}
	for rows.Next() {
		var i City
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Lat,
			&i.Long,
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