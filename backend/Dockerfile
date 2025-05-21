FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main /app/cmd/api/main.go

# Run stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migrations ./db/migrations

RUN chmod +x start.sh
RUN chmod +x wait-for.sh

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]