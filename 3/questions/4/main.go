package main

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment: filename=ascii_sample.zip")

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	filename := "./dest.txt"
	writer, err := zipWriter.Create(filename)
	if err != nil {
		panic(err)
	}

	io.Copy(writer, strings.NewReader("Nice Request!!!"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
