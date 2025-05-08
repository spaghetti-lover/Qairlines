package main

import (
	_ "github.com/jackc/pgx/v5"
	"github.com/spaghetti-lover/qairlines/utils"
)

func main() {
	print(utils.RandomStringNum())
}
