-- name: CreateTransfer :one
INSERT INTO transfers (from_account_id, to_account_id, amount)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetTransfer :one
SELECT *
FROM transfers
WHERE id = $1
LIMIT 1;
-- name: GetTransfers :many
SELECT *
FROM transfers
ORDER BY id
LIMIT $1 offset $2;
-- name: UpdateTransfer :one
update transfers
set amount = $2
where id = $1
returning *;
-- name: DeleteTransfer :exec
delete from transfers
where id = $1;