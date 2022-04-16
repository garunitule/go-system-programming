package main

import (
	"io"
	"os"
	"strings"
)

func copyN(w io.Writer, r io.Reader, length int64) (witten int64, err error) {
	return io.Copy(w, io.LimitReader(r, length))
}

func main() {
	copyN(os.Stdout, strings.NewReader("hogehoge fuafuga piyopiyo"), 10)
}
