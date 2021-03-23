-- name: CreateEntry :one
INSERT INTO entries (account_id, amount)
VALUES ($1, $2)
RETURNING *;
-- name: GetEntry :one
SELECT *
FROM entries
WHERE id = $1
LIMIT 1;
-- name: GetEntries :many
SELECT *
FROM entries
ORDER BY id
LIMIT $1 offset $2;
-- name: UpdateEntry :one
update entries
set amount = $2
where id = $1
returning *;
-- name: DeleteEntry :exec
delete from entries
where id = $1;