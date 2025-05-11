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

-- name: CountOccupiedSeats :one
SELECT COUNT(*) FROM passengers as p
JOIN booking as b ON p.booking_id = b.booking_id
WHERE b.flight_id = $1 AND b.flight_class = $2;

-- name: CheckSeatOccupied :one
SELECT EXISTS (
  SELECT 1
  FROM passengers
  JOIN booking ON passengers.booking_id = booking.booking_id
  WHERE booking.flight_id = $1
    AND booking.flight_class = $2
    AND passengers.seat_row = $3
    AND passengers.seat_col = $4
) AS seat_taken;

-- name: ListPassengers :many
SELECT * FROM passengers
ORDER BY passenger_id
LIMIT $1
OFFSET $2;


-- name: DeletePassenger :exec
DELETE FROM passengers
WHERE passenger_id = $1;