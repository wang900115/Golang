package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Block struct {
	PreviousHash []byte
	Hash         []byte
	Data         []byte
	Nonce        int
}

func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{previousHash, []byte{}, []byte(data), 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	fmt.Printf("Creating block with data: %s\n", data)
	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
