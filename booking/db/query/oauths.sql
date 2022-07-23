-- name: CreateOauths :one
insert into oauths (
    "user_id",
    "oauth_type",
    "oauth_id",
    "unionid",
    "credential"
) values (
    $1, $2, $3, $4, $5
) returning *;

-- name: UpdateOauths :one
update oauths set 
    user_id = $2,
    oauth_id = $3,
    unionid = $4,
    credential = $5
where id = $1 returning *;
 