package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index int64	// q区块编号
	Timestamp int64 // 时间戳
	PreBlockHash string	// 上一个区块哈希
	Hash string	// d当前区块哈希值

	Data string	//区块链数据
}

// 计算hash
func calculateHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PreBlockHash + b.Data
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

// 生成区块
func GenerateNewBlock(prevBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.PreBlockHash = prevBlock.Hash
	// 必须等data赋值完后再计算Hash
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

// 创世区块
func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.PreBlockHash = ""
	return GenerateNewBlock(preBlock, "Genesis Block")
}

