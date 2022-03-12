package main

import (
	"github.com/plusbeauxjours/GO-crypto/blockchain"
	"fmt"
)

func main() {
	chain := blockchain.GetBlockchain()
	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Prev Hash: %s\n", block.PrevHash)
	}
}
