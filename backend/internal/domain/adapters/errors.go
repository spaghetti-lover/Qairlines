package adapters

import "errors"

var (
	ErrTicketNotFound          = errors.New("ticket not found")
	ErrTicketCannotBeCancelled = errors.New("ticket cannot be cancelled due to its current status")
)
