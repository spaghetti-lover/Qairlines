package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spaghetti-lover/qairlines/config"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/handlers"
	"github.com/spaghetti-lover/qairlines/internal/infra/kafka"
	"github.com/spaghetti-lover/qairlines/internal/infra/mailer"
)

func main() {
	config, err := config.LoadConfig("../backend")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Tạo mailer
	smtpMailer := &mailer.SMTPMailer{
		From:     config.MailFrom,
		Password: config.MailPassword,
		Host:     config.MailHost,
		Port:     config.MailPort,
	}

	// Khởi tạo và chạy consumer
	consumer := kafka.NewMailConsumer(
		config.KafkaBrokerURL,
		config.KafkaTopic,
		config.KafkaGroupID,
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
		config.KafkaBrokerURL,
		config.KafkaTopic,
	)

	// Setup HTTP server
	mailRepo := mailer.NewMailRepository(producer)
	mailUseCase := usecases.NewMailUseCase(mailRepo)
	mailHandler := handlers.NewSendMailHandler(mailUseCase)

	mux := http.NewServeMux()
	mux.Handle("/send-mail", mailHandler)

	server := &http.Server{
		Addr:    config.MailServer,
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
