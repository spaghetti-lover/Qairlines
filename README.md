# Clean Architecture for dumb shit web in class

/qairline-backend
│
├── main.go
├── config/
├── models/
├── handlers/
├── services/
├── repositories/
├── routes/
├── utils/
├── db/
│ └── migrations/
│ └── mocks/
│ ├── sqlc/
│ └── queries/

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
