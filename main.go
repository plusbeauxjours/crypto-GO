package main

import "github.com/plusbeauxjours/GO-crypto/blockchain"

func main() {
	blockchain.Blockchain().AddBlock("First")
	blockchain.Blockchain().AddBlock("Second")
	blockchain.Blockchain().AddBlock("Third")
}
