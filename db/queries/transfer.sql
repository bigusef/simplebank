-- name: CreateTransfer :one
INSERT INTO transfers(from_account_id, to_account_id, amount)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetTransfer :one
SELECT *
FROM transfers
WHERE id = $1
LIMIT 1;

-- name: ListTransfer :many
SELECT *
FROM transfers
LIMIT $1 OFFSET $2;

-- name: UpdateTransfer :one
UPDATE transfers
SET from_account_id=$1,
    to_account_id=$2,
    amount=$3
WHERE id = $4
RETURNING *;

-- name: DeleteTransfer :exec
DELETE
FROM transfers
WHERE id = $1;