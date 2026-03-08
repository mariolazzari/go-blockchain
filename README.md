# Golang: How to Build a Blockchain in Go Guide

## Introduction

- Go lang
- VS code

## What is a blockchain

### Definition

Distributed data structure used to store records (transactions or events) in a sequence of blocks that are cryptographically linked and replicated across multiple nodes in a network.

#### Core definition

A blockchain is a decentralized, append-only ledger where data is grouped into blocks, each block referencing the cryptographic hash of the previous block, forming a chronological chain.

#### Node structure

- Transactions / data – the information being recorded
- Timestamp – time when the block was created
- Previous block hash – cryptographic reference to the prior block
- Nonce – a value used in consensus mechanisms (e.g., mining)
- Block hash – the hash of the block contents

### Typical use cases

- Cryptocurrencies (e.g., Bitcoin, Ethereum)
- Supply chain tracking
- Smart contracts
- Digital identity systems
- Audit trails and notarization

## Creating blockchain

### Project setup

```sh
go mod init github.com/mariolazzari/goblockchain
touch blockchain.go
air init
air
```

### Create a block

```go
package main

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	return &Block{
		timestamp:    time.Now().UnixNano(),
		nonce:        nonce,
		previousHash: previousHash,
	}
}

func (b *Block) Print() {
	fmt.Printf("timestamp       %d\n", b.timestamp)
	fmt.Printf("nonce           %d\n", b.nonce)
	fmt.Printf("previous_hash   %s\n", b.previousHash)
	fmt.Printf("transactions    %s\n", b.transactions)
}

func init() {
	log.SetPrefix("Bolckchian: ")
}

func main() {
	b := NewBlock(0, "init hash")
	b.Print()
}
```

### Blockchain struct

```go
package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	return &Block{
		timestamp:    time.Now().UnixNano(),
		nonce:        nonce,
		previousHash: previousHash,
	}
}

func (b *Block) Print() {
	fmt.Printf("timestamp       %d\n", b.timestamp)
	fmt.Printf("nonce           %d\n", b.nonce)
	fmt.Printf("previous_hash   %s\n", b.previousHash)
	fmt.Printf("transactions    %s\n", b.transactions)
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i,
			strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func init() {
	log.SetPrefix("Bolckchian: ")
}

func main() {
	blockChain := NewBlockchain()
	blockChain.Print()
	blockChain.CreateBlock(5, "hash 1")
	blockChain.Print()
	blockChain.CreateBlock(2, "hash 2")
	blockChain.Print()
}
```
