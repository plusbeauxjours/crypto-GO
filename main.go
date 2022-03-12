package main

import (
	"github.com/plusbeauxjours/GO-crypto/blockchain"
	"fmt"
)

func main() {
	chain := blockchain.GetBlockchain()
	fmt.Println(chain)
}
