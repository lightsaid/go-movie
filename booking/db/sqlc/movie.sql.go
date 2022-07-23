// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: movie.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createMovie = `-- name: CreateMovie :one
insert into movie (
    title,
    release_at,
    cover_url,
    duration,
    language,
    genre,
    rating,
    director,
    "desc",
    status,
    star,
    wish_count
)values(
    $1, $2, $3, $4, $5,
    $6, $7, $8, $9, $10,
    $11, $12
) returning id, title, release_at, cover_url, duration, language, genre, rating, director, "desc", status, star, wish_count
`

type CreateMovieParams struct {
	Title     string         `json:"title"`
	ReleaseAt time.Time      `json:"release_at"`
	CoverUrl  string         `json:"cover_url"`
	Duration  sql.NullTime   `json:"duration"`
	Language  sql.NullString `json:"language"`
	Genre     sql.NullString `json:"genre"`
	Rating    sql.NullString `json:"rating"`
	Director  sql.NullString `json:"director"`
	Desc      sql.NullString `json:"desc"`
	Status    int32          `json:"status"`
	Star      sql.NullString `json:"star"`
	WishCount sql.NullInt32  `json:"wish_count"`
}

func (q *Queries) CreateMovie(ctx context.Context, arg CreateMovieParams) (Movie, error) {
	row := q.db.QueryRowContext(ctx, createMovie,
		arg.Title,
		arg.ReleaseAt,
		arg.CoverUrl,
		arg.Duration,
		arg.Language,
		arg.Genre,
		arg.Rating,
		arg.Director,
		arg.Desc,
		arg.Status,
		arg.Star,
		arg.WishCount,
	)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ReleaseAt,
		&i.CoverUrl,
		&i.Duration,
		&i.Language,
		&i.Genre,
		&i.Rating,
		&i.Director,
		&i.Desc,
		&i.Status,
		&i.Star,
		&i.WishCount,
	)
	return i, err
}

const getMovie = `-- name: GetMovie :one
select id, title, release_at, cover_url, duration, language, genre, rating, director, "desc", status, star, wish_count from movie where id = $1 limit 1
`

func (q *Queries) GetMovie(ctx context.Context, id int64) (Movie, error) {
	row := q.db.QueryRowContext(ctx, getMovie, id)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ReleaseAt,
		&i.CoverUrl,
		&i.Duration,
		&i.Language,
		&i.Genre,
		&i.Rating,
		&i.Director,
		&i.Desc,
		&i.Status,
		&i.Star,
		&i.WishCount,
	)
	return i, err
}

const getMovies = `-- name: GetMovies :many
select id, title, release_at, cover_url, duration, language, genre, rating, director, "desc", status, star, wish_count from movie limit $1 offset $2
`

type GetMoviesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetMovies(ctx context.Context, arg GetMoviesParams) ([]Movie, error) {
	rows, err := q.db.QueryContext(ctx, getMovies, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Movie{}
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.ReleaseAt,
			&i.CoverUrl,
			&i.Duration,
			&i.Language,
			&i.Genre,
			&i.Rating,
			&i.Director,
			&i.Desc,
			&i.Status,
			&i.Star,
			&i.WishCount,
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

const getMoviesByStatus = `-- name: GetMoviesByStatus :many
select id, title, release_at, cover_url, duration, language, genre, rating, director, "desc", status, star, wish_count from movie where status = $1 limit $2 offset $3
`

type GetMoviesByStatusParams struct {
	Status int32 `json:"status"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetMoviesByStatus(ctx context.Context, arg GetMoviesByStatusParams) ([]Movie, error) {
	rows, err := q.db.QueryContext(ctx, getMoviesByStatus, arg.Status, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Movie{}
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.ReleaseAt,
			&i.CoverUrl,
			&i.Duration,
			&i.Language,
			&i.Genre,
			&i.Rating,
			&i.Director,
			&i.Desc,
			&i.Status,
			&i.Star,
			&i.WishCount,
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

const updateMoive = `-- name: UpdateMoive :one
update movie set 
    title=$2,
    release_at=$3,
    cover_url=$4,
    duration=$5,
    language=$6,
    genre=$7,
    rating=$8,
    director=$9,
    "desc"=$10,
    status=$11,
    star=$12,
    wish_count=$13
where id = $1 returning id, title, release_at, cover_url, duration, language, genre, rating, director, "desc", status, star, wish_count
`

type UpdateMoiveParams struct {
	ID        int64          `json:"id"`
	Title     string         `json:"title"`
	ReleaseAt time.Time      `json:"release_at"`
	CoverUrl  string         `json:"cover_url"`
	Duration  sql.NullTime   `json:"duration"`
	Language  sql.NullString `json:"language"`
	Genre     sql.NullString `json:"genre"`
	Rating    sql.NullString `json:"rating"`
	Director  sql.NullString `json:"director"`
	Desc      sql.NullString `json:"desc"`
	Status    int32          `json:"status"`
	Star      sql.NullString `json:"star"`
	WishCount sql.NullInt32  `json:"wish_count"`
}

func (q *Queries) UpdateMoive(ctx context.Context, arg UpdateMoiveParams) (Movie, error) {
	row := q.db.QueryRowContext(ctx, updateMoive,
		arg.ID,
		arg.Title,
		arg.ReleaseAt,
		arg.CoverUrl,
		arg.Duration,
		arg.Language,
		arg.Genre,
		arg.Rating,
		arg.Director,
		arg.Desc,
		arg.Status,
		arg.Star,
		arg.WishCount,
	)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ReleaseAt,
		&i.CoverUrl,
		&i.Duration,
		&i.Language,
		&i.Genre,
		&i.Rating,
		&i.Director,
		&i.Desc,
		&i.Status,
		&i.Star,
		&i.WishCount,
	)
	return i, err
}