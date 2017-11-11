/**
 * Learning from https://github.com/Jeiwan/blockchain_go
 */

package main

import (
	"fmt"

	"github.com/furusiyya/blockchain/core"
)

func main() {
	bc := core.NewBlockchain()
	bc.AddBlock("Send 1 BTC to Bilal")
	bc.AddBlock("Send 2 BTC to Danish")
	bc.AddBlock("Send 3 BTC to Everyone")

	for _, block := range bc.Blocks {
		fmt.Println(block.String())
	}
}
