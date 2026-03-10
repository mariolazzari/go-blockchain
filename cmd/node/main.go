package main

import (
	"log"

	"github.com/mariolazzari/goblockchain/internal/blockchain"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockChain := blockchain.NewBlockchain()
	blockChain.Print()

	blockChain.AddTransaction("A", "B", 1.0)
	previousHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(5, previousHash)
	blockChain.Print()

	blockChain.AddTransaction("C", "D", 2.0)
	blockChain.AddTransaction("X", "Y", 3.0)
	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(2, previousHash)
	blockChain.Print()
}
