package core

import ()

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func (bs *Blockchain) AddBlock(data string) {
	newBlock := NewBlock(data, bs.getLatestBlockHash())
	bs.Blocks = append(bs.Blocks, newBlock)
}

func (bs *Blockchain) getLatestBlockHash() []byte {
	return bs.Blocks[len(bs.Blocks)-1].hash
}
