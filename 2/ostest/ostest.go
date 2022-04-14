package ostest

import (
	"io"
	"os"
)

func WriteToFile() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File example\n"))
	file.Close()
}

func WriteToScreen() {
	os.Stdout.Write([]byte("os.Stdout example\n"))
}

func MultiWriter() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")
}
