package main

import (
	"fmt"

	_ "github.com/jackc/pgx/v5"
)

func main() {
	fmt.Println(string('A' + 1))
}
