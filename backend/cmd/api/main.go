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
	"github.com/redis/go-redis/v9"
	"github.com/spaghetti-lover/qairlines/config"
	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/infra/api"
	"github.com/spaghetti-lover/qairlines/internal/infra/mail"
	"github.com/spaghetti-lover/qairlines/internal/infra/worker"
	"github.com/spaghetti-lover/qairlines/pkg/logger"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
	"golang.org/x/sync/errgroup"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	err = utils.LoadMessages("./message.json")
	if err != nil {
		log.Fatalf("Failed to load messages: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	poolConfig, err := pgxpool.ParseConfig(cfg.DBSource)
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
	redis := config.NewRedisClient()
	if redis == nil {
		log.Fatal("failed to create Redis client")
	}
	store := db.NewStore(connPool)

	redisOpt := asynq.RedisClientOpt{
		Addr: cfg.RedisAddress,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)

	waitGroup, ctx := errgroup.WithContext(ctx)

	// Start task processor in goroutine
	runTaskProcessor(ctx, waitGroup, cfg, redisOpt, store)
	// Start server in goroutine
	runApiServer(ctx, waitGroup, cfg, redis, store, taskDistributor)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal("Error in goroutines: ", err)
	}
}

func runApiServer(ctx context.Context, waitGroup *errgroup.Group, config config.Config, redis *redis.Client, store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := api.NewServer(config, store, redis, taskDistributor)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	httpServer := &http.Server{
		Addr:    config.ServerAddressPort,
		Handler: server,
	}

	waitGroup.Go(func() error {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe failed: %v", err)
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Println("Shutting down server...")
		err := httpServer.Shutdown(context.Background())
		if err != nil {
			log.Fatalf("Server shutdown failed: %v", err)
			return err
		}
		log.Println("Server gracefully stopped")
		return nil
	})
}

func runTaskProcessor(ctx context.Context, waitGroup *errgroup.Group, config config.Config, redisOpt asynq.RedisClientOpt, store db.Store) {
	mailer := mail.NewGmailSender(config.MailSenderName, config.MailSenderAddress, config.MailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)
	log.Println("Task processor started")
	if err := taskProcessor.Start(); err != nil {
		log.Fatalf("Failed to start task processor: %v", err)
	}
	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Println("Shutting down task processor...")
		taskProcessor.Shutdown()
		log.Println("Task processor gracefully stopped")
		return nil
	})
}
