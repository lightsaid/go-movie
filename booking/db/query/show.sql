-- name: CreateShow :one
insert into show (
    "date",
    "start_time",
    "end_time",
    "cinema_hall_id",
    "movie_id"
) values (
    $1, $2, $3, $4, $5
) returning *;

-- name: GetShow :one
select * from show where id = $1;

-- name: GetShowList :many
select * from show limit $1 offset $2;

-- name: UpdateShow :one 
update show set
    date = $2, 
    start_time = $3,
    end_time = $4,
    cinema_hall_id = $5,
    movie_id = $6
where id = $1 returning *; 

