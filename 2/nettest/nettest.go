package nettest

import (
	"io"
	"net"
	"net/http"
	"os"
)

func Write() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	io.Copy(os.Stdout, conn)
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWriter sample")
}

func WebServe() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
