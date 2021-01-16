package main

import (
	"fmt"
	"github.com/PatrykZdan/Blockchain-golang/blockchain"
	"strconv"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		proof := blockchain.NewProof(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(proof.Validate()))
		fmt.Println()
	}
}
