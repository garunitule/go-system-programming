package fileread

import (
	"io"
	"os"
)

func Read() {
	file, err := os.Open("/Users/kogayuuta/study/go-syspro/3/file.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
