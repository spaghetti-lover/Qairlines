package main

import (
	"fmt"

	_ "github.com/jackc/pgx/v5"
	"github.com/spaghetti-lover/qairlines/utils"
)

func main() {
	fmt.Print(utils.RandomEmail())
}
