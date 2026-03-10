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
package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mariolazzari/goblockchain/internal/tx"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*tx.Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*tx.Transaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp       %d\n", b.timestamp)
	fmt.Printf("nonce           %d\n", b.nonce)
	fmt.Printf("previous_hash   %x\n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64             `json:"timestamp"`
		Nonce        int               `json:"nonce"`
		PreviousHash [32]byte          `json:"previous_hash"`
		Transactions []*tx.Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}
```

### Blockchain struct

```go
package blockchain

import (
	"fmt"
	"strings"

	"github.com/mariolazzari/goblockchain/internal/tx"
)

type Blockchain struct {
	transactionPool []*tx.Transaction
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*tx.Transaction{}
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i,
			strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := tx.NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}
```

### Compute hash

[Go package](https://pkg.go.dev/crypto/sha256)

### Add a transaction

```go
package tx

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address      %s\n", t.senderBlockchainAddress)
	fmt.Printf(" recipient_blockchain_address   %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value                          %.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}
```

### 
