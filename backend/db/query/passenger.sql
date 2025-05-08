-- name: CreatePassenger :one
INSERT INTO passengers (
  booking_id,
  citizen_id,
  passport_number,
  gender,
  phone_number,
  first_name,
  last_name,
  nationality,
  date_of_birth,
  seat_row,
  seat_col
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;


-- name: GetPassenger :one
SELECT * FROM passengers
WHERE passenger_id = $1 LIMIT 1;

-- name: ListPassengers :many
SELECT * FROM passengers
ORDER BY passenger_id
LIMIT $1
OFFSET $2;


-- name: DeletePassenger :exec
DELETE FROM passengers
WHERE passenger_id = $1;