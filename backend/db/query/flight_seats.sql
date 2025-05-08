-- name: CreateFlightSeat :one
INSERT INTO flight_seats (
  registration_number,
  flight_class,
  class_multiplier,
  child_multiplier,
  max_row_seat,
  max_col_seat
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetFlightSeat :one
SELECT * FROM flight_seats
WHERE registration_number = $1 LIMIT 1;

-- name: ListFlightSeats :many
SELECT * FROM flight_seats
ORDER BY registration_number
LIMIT $1
OFFSET $2;


-- name: DeleteFlightSeat :exec
DELETE FROM flight_seats
WHERE registration_number = $1;