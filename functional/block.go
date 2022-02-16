package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
)

func GenerateBlock() string {
	transactions := []string{"Transaction 1", "Transaction 2", "Transaction 3", "Transaction 4"}

	return fmt.Sprintf("[blockid= %d, Transactions = %s, nonce = %d, difficulty = %d, preBlockHash = %s]",
		blockId, transactions, nonce, difficulty, preBlockHash)
}

func HashBlock(block string) string {
	blockHashBytes := sha256.Sum256([]byte(block))
	return fmt.Sprintf("%x", blockHashBytes)
}

func IsValidBlock(blockHash string) bool {
	isValid := strings.HasPrefix(blockHash, difficultyStr)
	return isValid
}
func SaveBlock(block string, blockHash string) {
	fmt.Println("Block :", block)
	fmt.Println("Block Hash:", blockHash)
	printSeperator()

	blockId++
	nonce = 0
	preBlockHash = blockHash
	optimizeDifficulty()
	lastBlockDate = time.Now()
}
