package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func startMine(cancelCh chan bool) {
	go hanldeOsSignal(cancelCh)

	select {
	case <-cancelCh:
		return
	default:
		for {
			// Step 1: Create Block
			block := GenerateBlock()

			// Step 2: Calculate Block Hash
			blockHash := HashBlock(block)

			//Step 3: Validate Block(Mine Block :))
			if IsValidBlock(blockHash) {
				SaveBlock(block, blockHash)
			} else {
				nonce++
			}
		}
	}
}

func hanldeOsSignal(cancelCh chan bool) {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)

	<-sigCh

	cancelCh <- true
}

func optimizeDifficulty() {
	dateDiff := time.Since(lastBlockDate)

	if dateDiff.Seconds() < 20 {
		difficulty++
	} else {
		difficulty--
	}

	if difficulty <= 0 {
		difficulty = 1
	}

	difficultyStr = strings.Repeat("0", difficulty)
	fmt.Printf("difficulty Changed to %d\r\n", difficulty)
}
