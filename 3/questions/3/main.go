package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	filename := "./dest.txt"
	zippedPackageName := os.Getenv("GO_SYSPRO_PATH") + "/3/questions/3/dest.zip"

	zippedPackage, err := os.Create(zippedPackageName)
	if err != nil {
		panic(err)
	}

	zipWriter := zip.NewWriter(zippedPackage)
	defer zipWriter.Close()

	writer, err := zipWriter.Create(filename)
	if err != nil {
		panic(err)
	}

	io.Copy(writer, strings.NewReader("test hogehoge\ntest hogehoge2"))
}
