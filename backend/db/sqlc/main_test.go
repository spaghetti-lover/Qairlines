package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/qairline?sslmode=disable"
)

var testStore Store

func TestMain(m *testing.M) {
	connPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Khong connect db duoc dcm, quay ve main_test xem di: ", err)
	}
	testStore = NewStore(connPool)
	//Stop testing and info test success or not
	os.Exit(m.Run())
}
