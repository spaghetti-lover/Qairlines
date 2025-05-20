package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/handlers"
	"github.com/spaghetti-lover/qairlines/internal/infra/kafka"
	"github.com/spaghetti-lover/qairlines/internal/infra/mailer"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Tạo mailer
	smtpMailer := &mailer.SMTPMailer{
		From:     os.Getenv("MAIL_FROM"),
		Password: os.Getenv("MAIL_PASSWORD"),
		Host:     os.Getenv("MAIL_HOST"),
		Port:     os.Getenv("MAIL_PORT"),
	}

	// Khởi tạo và chạy consumer
	consumer := kafka.NewMailConsumer(
		os.Getenv("KAFKA_BROKER_URL"),
		os.Getenv("KAFKA_TOPIC"),
		os.Getenv("KAFKA_GROUP_ID"),
		smtpMailer,
	)

	go func() {
		log.Println("Starting mail consumer...")
		if err := consumer.Start(ctx); err != nil {
			log.Fatalf("Mail consumer error: %v", err)
		}
	}()

	// Khởi tạo producer và inject vào handler
	producer := kafka.NewMailProducer(
		os.Getenv("KAFKA_BROKER"),
		os.Getenv("KAFKA_TOPIC"),
	)

	// Setup HTTP server
	mailRepo := mailer.NewMailRepository(producer)
	mailUseCase := usecases.NewMailUseCase(mailRepo)
	mailHandler := handlers.NewSendMailHandler(mailUseCase)

	mux := http.NewServeMux()
	mux.Handle("/send-mail", mailHandler)

	server := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	// Start HTTP server in a goroutine
	go func() {
		log.Println("Server started at :8081")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-ctx.Done()
	log.Println("Shutting down mail service...")

	// Create shutdown context with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown both consumer and HTTP server gracefully
	if err := consumer.Close(); err != nil {
		log.Fatalf("Consumer shutdown failed: %v", err)
	}

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Mail server gracefully stopped")
}
