-- name: CreateFlightSeat :one
INSERT INTO flight_seats (
  flight_id,
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
WHERE flight_id = $1 AND flight_class = $2;

-- name: ListFlightSeats :many
SELECT * FROM flight_seats
ORDER BY flight_id
LIMIT $1
OFFSET $2;

-- name: ListFlightSeatsByFlightID :many
SELECT * FROM flight_seats
WHERE flight_id = $1;


-- name: DeleteFlightSeat :exec
DELETE FROM flight_seats
WHERE flight_id = $1;