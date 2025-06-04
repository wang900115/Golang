package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	PreviousHash []byte
	Hash         []byte
	Data         []byte
}

type Blockchain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	hash := sha256.Sum256(bytes.Join([][]byte{b.Data, b.PreviousHash}, []byte{}))
	b.Hash = hash[:]
}

func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{previousHash, []byte(data), []byte{}}
	block.DeriveHash()
	return block
}

func (chain *Blockchain) AddBlock(data string) {
	previousHash := chain.blocks[len(chain.blocks)-1].Hash
	newBlock := CreateBlock(data, previousHash)
	chain.blocks = append(chain.blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()
	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.blocks {
		fmt.Printf("Prev Hash: %x \n", block.PreviousHash)
		fmt.Printf("Data: %s \n", block.Data)
		fmt.Printf("Hash: %x \n", block.Hash)
	}

}
