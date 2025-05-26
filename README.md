# Clean Architecture for dumb shit web in class

```text
qairline-backend/
├── go.mod
├── go.sum
├── README.md

├── cmd/
│   ├── api/
│   │   └── main.go         # Entry point for HTTP server
│   └── migrate/
│       └── main.go         # Entry point for DB migration CLI

├── config/
│   └── config.go           # Config loading (env, flags, etc.)

├── internal/
│   └── api/
│       ├── handler.go  # HTTP handler wiring
│       └── booking.go  # HTTP endpoints for bookings


├── domain/
│   └── booking/
│       ├── service.go         # Core business logic
│       ├── service_test.go
│       ├── model.go           # Booking model
│       ├── postgres/
│       │   └── storage.go     # PostgreSQL-specific implementation
│       └── mock/
│           └── service.go     # Mock implementation for testing

├── pkg/
│   └── utils/
│       └── logger.go          # Logging utility

├── db/
│   ├── migrations/
│   │   └── 001_init.sql
│   └── sqlc/
│       ├── queries.sql        # SQL queries
│       └── models.go          # Generated models
```

## Reason for techstack:

### SQLC

| Feature            | Database/SQL                                    | SQLx                                        | GORM                                                   | SQLC                                           |
| ------------------ | ----------------------------------------------- | ------------------------------------------- | ------------------------------------------------------ | ---------------------------------------------- |
| Speed & Ease       | Very fast & straightforward                     | Quite fast & easy to use                    | CRUD functions implemented, very short production code | Very fast & easy to use                        |
| Field Mapping      | Manual mapping SQL fields to variables          | Fields mapping via query text & struct tags | Must learn to write queries using gorm's function      | Automatic code generation                      |
| Error Detection    | Easy to make mistakes, not caught until runtime | Failure won't occur until runtime           | Must learn to write queries using gorm's function      | Catch SQL query errors before generating codes |
| Performance (Load) | -                                               | -                                           | Run slowly on high load                                | -                                              |

### Take note:

- Tích hợp MinIO với backend, dùng cho việc upload, download file
- Hoan thien transaction booking voi concurrent request
- Test chuc nang sending mail va test performance sao lai cham vay
- Implement not phan BookingHistory cho cac /api/user