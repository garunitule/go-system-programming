package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	sendMessages := []string{"hello", "world", "!!"}
	current := 0
	var conn net.Conn
	for {
		var err error
		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8888")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access %v¥n", current)
		}

		request, err := http.NewRequest("POST", "http://localhost:8888", strings.NewReader(sendMessages[current]))
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
		// タイムアウト時の挙動確認
		// time.Sleep(time.Second * 5)

		current++
		if current == len(sendMessages) {
			break
		}
	}
	conn.Close()
}
