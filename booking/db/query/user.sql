-- name: CreateUser :one
insert into "user" (
    name,
    avatar_url,
    password,
    phone
) values (
    $1, $2, $3, $4
) returning *;

-- name: GetUser :one 
select * from "user" where id = $1; 

-- name: UpdateUser :one
update "user" set 
    name = $2,
    avatar_url = $3,
    password = $4,
    phone = $5
where id = $1 returning *;


