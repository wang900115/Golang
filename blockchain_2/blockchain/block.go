package blockchain

type Block struct {
	PreviousHash []byte
	Hash         []byte
	Data         []byte
	Nonce        int
}

type Blockchain struct {
	Blocks []*Block
}

func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{previousHash, []byte(data), []byte{}, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (chain *Blockchain) AddBlock(data string) {
	previousHash := chain.Blocks[len(chain.Blocks)-1].Hash
	newBlock := CreateBlock(data, previousHash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}
