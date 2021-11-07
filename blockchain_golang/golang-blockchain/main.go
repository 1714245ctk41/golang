package main

import (
	"blockchain-golang/blockchain"
	"fmt"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")
	for _, block := range chain.Blocks {
		fmt.Printf("Previeous Hash: %x\n", block.PrevHas)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("___________")
	}
}
