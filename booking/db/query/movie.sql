-- name: CreateMovie :one
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
) returning *;

-- name: UpdateMoive :one 
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
where id = $1 returning *;

-- name: GetMovie :one
select * from movie where id = $1 limit 1;

-- name: GetMoviesByStatus :many
select * from movie where status = $1 limit $2 offset $3;

-- name: GetMovies :many
select * from movie limit $1 offset $2;

