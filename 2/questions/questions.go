package main

import (
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	answer1()
	answer2()
	answer3()
}

func answer1() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "Write to File %d\n", 1)
	fmt.Fprintf(file, "Write to File %s\n", "string")
	fmt.Fprintf(file, "Write to File %f\n", 0.001)
	file.Close()
}

func answer2() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records)

	if err := w.Error(); err != nil {
		panic(err)
	}
}

func handler(serverWriter http.ResponseWriter, r *http.Request) {
	serverWriter.Header().Set("Content-Encoding", "gzip")
	serverWriter.Header().Set("Content-Type", "application/json")
	gzipWriter := gzip.NewWriter(serverWriter)

	source := map[string]string{
		"Hello": "World",
	}

	multiWriter := io.MultiWriter(gzipWriter, os.Stdout)
	serverEncoder := json.NewEncoder(multiWriter)
	serverEncoder.Encode(source)
	gzipWriter.Flush()
}

func answer3() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
