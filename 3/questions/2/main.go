package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile(os.Getenv("GO_SYSPRO_PATH")+"/3/questions/2/new.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	n, err := io.Copy(file, io.LimitReader(rand.Reader, 1024))
	if err != nil {
		panic(err)
	}
	fmt.Printf("witten length: %d\n", n)
}
