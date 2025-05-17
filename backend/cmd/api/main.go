package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"net/http"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	api "github.com/spaghetti-lover/qairlines/internal/api"
)

const (
	dbDriver      = ""
	dbSource      = ""
	serverAddress = ":8000"
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

	store := db.NewStore(connPool)
	server, err := api.NewServer(&store)
	if err != nil {
		log.Fatal("cannot connect to server: ", err)
	}
	log.Println("Server start on ", serverAddress)
	go func() {
		err = http.ListenAndServe(":8000", server)
		if err != nil {
			log.Fatal("cannot start server: ", err)
		}
	}()
	<-ctx.Done()
	log.Println("Shutting down gracefully...")
}
