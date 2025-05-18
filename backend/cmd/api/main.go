package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/infra/api"
)

const (
	dbSource      = "postgresql://root:secret@localhost:5432/qairline?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	connPool, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}
	defer connPool.Close()

	store := db.NewStore(connPool)

	// Create server
	server, err := api.NewServer(&store)
	if err != nil {
		log.Fatal("cannot connect to server: ", err)
	}

	// Setup HTTP server
	httpServer := &http.Server{
		Addr:    serverAddress,
		Handler: server,
	}

	// Start server in goroutine
	go func() {
		log.Println("Server started on", serverAddress)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe failed: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server gracefully stopped")
}
