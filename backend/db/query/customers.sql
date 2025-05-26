-- name: CreateCustomer :one
INSERT INTO customers (
  user_id,
  phone_number,
  gender,
  date_of_birth,
  passport_number,
  identification_number,
  address,
  loyalty_points
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetCustomer :one
SELECT * FROM customers
WHERE user_id = $1 LIMIT 1;

-- name: GetCustomerByEmail :one
SELECT c.* FROM customers c
JOIN users u ON c.user_id = u.user_id
WHERE u.email = $1 LIMIT 1;


-- name: GetAllCustomer :many
SELECT * FROM customers;

-- name: UpdateCustomer :exec
UPDATE customers
SET
  phone_number = $1,
  gender = $2,
  date_of_birth = $3,
  passport_number = $4,
  identification_number = $5,
  address = $6,
  loyalty_points = $7,
  updated_at = NOW()
WHERE user_id = $8;


-- name: ListCustomers :many
SELECT * FROM customers
ORDER BY user_id
LIMIT $1
OFFSET $2;


-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE user_id = $1;