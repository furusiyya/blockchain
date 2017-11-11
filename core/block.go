package core

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	timestamp     int64
	data          []byte
	prevBlockHash []byte
	hash          []byte
}

func GenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		timestamp:     time.Now().UnixNano(),
		data:          []byte(data),
		prevBlockHash: prevBlockHash,
	}
	block.setHash()
	return block
}

func (b *Block) setHash() {
	timestamp := []byte(strconv.FormatInt(b.timestamp, 10))
	headers := bytes.Join([][]byte{b.prevBlockHash, b.data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.hash = hash[:]
}

func (b *Block) String() string {
	return fmt.Sprintf("Prev. hash: %x\n Data: %s\n Hash: %x\n",
		b.prevBlockHash, b.data, b.hash)
}
