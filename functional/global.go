package main

import (
	"strings"
	"time"
)

var blockId int = 1
var difficulty int = 1
var difficultyStr string = "0"
var lastBlockDate time.Time
var preBlockHash string = ""
var nonce int = 0
var startSep string = strings.Repeat("*", 75)
