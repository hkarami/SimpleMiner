package main

import "fmt"

func printSeperator() {
	fmt.Println(startSep)
}

func loadConfig() {
	printSeperator()

	fmt.Println("Loading...............")

	fmt.Println("\t Load Settings.....")
	fmt.Println("\t Config Storage.....")
	fmt.Println("\t Config Network.....")
	fmt.Println("\t Config Nodes.....")
	fmt.Println("\t .................")
	fmt.Println("\t .................")
	fmt.Println("................Loaded")

	printSeperator()
}
