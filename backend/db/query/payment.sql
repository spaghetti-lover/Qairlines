-- name: CreatePayment :one
INSERT INTO payment (
  amount,
  currency,
  payment_method,
  status,
  booking_id
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;


-- name: GetPayment :one
SELECT * FROM payment
WHERE payment_id = $1 LIMIT 1;

-- name: ListPayment :many
SELECT * FROM payment
ORDER BY payment_id
LIMIT $1
OFFSET $2;


-- name: DeletePayment :exec
DELETE FROM payment
WHERE payment_id = $1;