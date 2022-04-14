package main

import (
	"./binaryread"
	"./bufferread"
)

func main() {
	// stdin.Read()
	// fileread.Read()
	// netread.ReadEasy()
	bufferread.Read()
	binaryread.Read()
	binaryread.ReadChangeEndian()
}
