package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func open() {
	file, err := os.Create("textfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "New file content\n")
}

func read() {
	file, err := os.Open("textfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("Read file")
	io.Copy(os.Stdout, file)
}

func throughputtest() {
	f, _ := os.Create("file.txt")
	a := time.Now()
	f.Write([]byte("緑の怪獣"))
	b := time.Now()
	f.Sync()
	c := time.Now()
	f.Close()
	d := time.Now()
	fmt.Printf("Write: %v\n", b.Sub(a))
	fmt.Printf("Sync: %v\n", c.Sub(b))
	fmt.Printf("Close: %v\n", d.Sub(c))
}

func main() {
	throughputtest()
	os.RemoveAll("test")
}
