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
  )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;
-- name: GetCustomer :one
SELECT *
FROM customers
WHERE user_id = $1
LIMIT 1;
-- name: GetCustomerByID :one
SELECT u.user_id AS uid,
  u.first_name,
  u.last_name,
  u.email,
  c.phone_number,
  c.date_of_birth,
  c.gender,
  c.identification_number,
  c.passport_number,
  c.address,
  c.loyalty_points
FROM Users u
  JOIN Customers c ON u.user_id = c.user_id
WHERE u.user_id = $1;
-- name: GetCustomerByEmail :one
SELECT c.*
FROM customers c
  JOIN users u ON c.user_id = u.user_id
WHERE u.email = $1
LIMIT 1;

-- name: UpdateCustomer :exec
UPDATE customers
SET phone_number = $1,
  gender = $2,
  date_of_birth = $3,
  passport_number = $4,
  identification_number = $5,
  address = $6,
  loyalty_points = $7
WHERE user_id = $8;

-- name: ListCustomers :many
SELECT *
FROM customers
ORDER BY user_id DESC
LIMIT $1 OFFSET $2;

-- name: DeleteCustomerByID :one
DELETE FROM Customers
WHERE user_id = $1
RETURNING user_id;