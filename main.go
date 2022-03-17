package main

import (
	"github.com/plusbeauxjours/GO-crypto/cli"
	"github.com/plusbeauxjours/GO-crypto/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
