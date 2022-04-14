package main

import (
	"fmt"
	"os"
	"time"

	"./jsontest"
)

type Talker interface {
	Talk()
}

type Greeter struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {
	var talker Talker
	talker = &Greeter{"wozozo"}
	talker.Talk()

	// ostest.WriteToFile()
	// ostest.WriteToScreen()
	// ostest.MultiWriter()
	// bytestest.WriteToBuffer()
	// buildertest.WriteToStringsBuffer()
	// nettest.Write()
	// nettest.WebServe()
	stdout := os.Stdout
	fmt.Fprintf(stdout, "Write with os.Stdout at %v", time.Now())
	jsontest.Write()
}
