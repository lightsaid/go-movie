// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
insert into "user" (
    name,
    avatar_url,
    password,
    phone
) values (
    $1, $2, $3, $4
) returning id, name, avatar_url, password, phone, created_at, updated_at
`

type CreateUserParams struct {
	Name      string         `json:"name"`
	AvatarUrl sql.NullString `json:"avatar_url"`
	Password  sql.NullString `json:"password"`
	Phone     sql.NullString `json:"phone"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Name,
		arg.AvatarUrl,
		arg.Password,
		arg.Phone,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AvatarUrl,
		&i.Password,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
select id, name, avatar_url, password, phone, created_at, updated_at from "user" where id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AvatarUrl,
		&i.Password,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
update "user" set 
    name = $2,
    avatar_url = $3,
    password = $4,
    phone = $5
where id = $1 returning id, name, avatar_url, password, phone, created_at, updated_at
`

type UpdateUserParams struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	AvatarUrl sql.NullString `json:"avatar_url"`
	Password  sql.NullString `json:"password"`
	Phone     sql.NullString `json:"phone"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Name,
		arg.AvatarUrl,
		arg.Password,
		arg.Phone,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AvatarUrl,
		&i.Password,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
