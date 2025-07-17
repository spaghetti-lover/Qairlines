package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"

	"github.com/hibiken/asynq"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/spaghetti-lover/qairlines/config"
	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/infra/api"
	"github.com/spaghetti-lover/qairlines/internal/infra/mail"
	"github.com/spaghetti-lover/qairlines/internal/infra/worker"
	"github.com/spaghetti-lover/qairlines/pkg/logger"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	err = utils.LoadMessages("./message.json")
	if err != nil {
		log.Fatalf("Failed to load messages: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	poolConfig, err := pgxpool.ParseConfig(config.DBSource)
	if err != nil {
		log.Fatal("cannot parse pool config:", err)
	}

	sqlLogger := logger.NewLoggerWithPath("logs/sql.log", "info")

	poolConfig.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger: &logger.PgxZerologTracer{
			Logger:         *sqlLogger,
			SlowQueryLimit: 500 * time.Millisecond,
		},
		LogLevel: tracelog.LogLevelDebug,
	}

	poolConfig.MaxConns = 20
	poolConfig.MinConns = 5
	poolConfig.MaxConnLifetime = 30 * time.Minute
	poolConfig.HealthCheckPeriod = 1 * time.Minute

	connPool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	defer connPool.Close()

	store := db.NewStore(connPool)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)

	// Create server
	server, err := api.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal("cannot connect to server: ", err)
	}

	// Setup HTTP server
	httpServer := &http.Server{
		Addr:    config.ServerAddressPort,
		Handler: server,
	}

	// Start task processor in goroutine
	go runTaskProcessor(config, redisOpt, store)
	// Start server in goroutine
	go runApiServer(config, store, taskDistributor)

	<-ctx.Done()
	log.Println("Shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Graceful shutdown for HTTP server
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server gracefully stopped")
}

func runApiServer(config config.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := api.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	httpServer := &http.Server{
		Addr:    config.ServerAddressPort,
		Handler: server,
	}

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe failed: %v", err)
	}
}

func runTaskProcessor(config config.Config, redisOpt asynq.RedisClientOpt, store db.Store) {
	mailer := mail.NewGmailSender(config.MailSenderName, config.MailSenderAddress, config.MailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)
	log.Println("Task processor started")
	if err := taskProcessor.Start(); err != nil {
		log.Fatalf("Failed to start task processor: %v", err)
	}
}
