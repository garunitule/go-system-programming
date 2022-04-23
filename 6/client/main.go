package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("GET", "http://localhost:8888", nil)
	if err != nil {
		panic(err)
	}

	err = request.Write(conn)
	if err != nil {
		panic(err)
	}

	response, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
}
