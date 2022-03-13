package main

import (
	"github.com/plusbeauxjours/GO-crypto/explorer"
	"github.com/plusbeauxjours/GO-crypto/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
