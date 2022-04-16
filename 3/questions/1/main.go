package main

import (
	"io"
	"os"
)

func main() {
	fileOld, err := os.Open("/Users/kogayuuta/study/go-syspro/3/questions/1/old.txt")
	if err != nil {
		panic(err)
	}

	fileNew, err := os.OpenFile("/Users/kogayuuta/study/go-syspro/3/questions/1/new.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	io.Copy(fileNew, fileOld)
}
