package main

import (
	"./binaryread"
	"./bufferread"
	"./textread"
)

func main() {
	// stdin.Read()
	// fileread.Read()
	// netread.ReadEasy()
	bufferread.Read()
	binaryread.Read()
	binaryread.ReadChangeEndian()
	textread.Read()
}
