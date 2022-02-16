package main

import (
	"fmt"
	"os"
)

func init() {
	loadConfig()
}

func main() {
	cancelCh := make(chan bool)
	go startMine(cancelCh)

	<-cancelCh

	fmt.Print("\r\n Exit")
	os.Exit(0)
}
