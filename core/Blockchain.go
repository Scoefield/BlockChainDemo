package core

import (
	"fmt"
	"log"
)

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := GenerateGenesisBlock()
	blockchain := Blockchain{}
	blockchain.AppendBlock(&genesisBlock)
	return &blockchain
}

func (bc *Blockchain) SendData(data string)  {
	prevBlock := bc.Blocks[len(bc.Blocks) - 1]
	newBlock := GenerateNewBlock(*prevBlock, data)
	bc.AppendBlock(&newBlock)
}

// 定义Blockchain的方法 AppendBlock：添加block
func (bc *Blockchain) AppendBlock(newBlock *Block)  {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks) - 1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid block!")
	}
}

func (bc *Blockchain) Print()  {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("prevHash: %s\n", block.PreBlockHash)
		fmt.Printf("currHash: %s\n", block.Hash)
		fmt.Printf("timestamp: %d\n", block.Timestamp)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Println()
	}
}

// 校验block
func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index - 1 != oldBlock.Index {
		fmt.Println("newBlock.Index - 1 != oldBlock.Index")
		return false
	}
	if newBlock.PreBlockHash != oldBlock.Hash {
		fmt.Println("newBlock.PreBlockHash != oldBlock.Hash")
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		fmt.Println("newBlock.PreBlockHash != oldBlock.Hash")
		return false
	}
	return true
}