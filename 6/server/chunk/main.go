package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

var contents = []string{
	"君も今たといその話院において事のうちへ見えるたた。",
	"もう多数から教育士はとこう大した安心たただけをしているたには指導寄ったでて、",
	"どうには忘れるただっますまし。",
	"哲学に致しましものは無論前をもっともたませな。",
	"何でもかでも嘉納さんに永続俗人あまり尊敬に思いです主",
	"義この現象あなたか増減がという実相談らしくでましますと、",
	"その毎号はそれか鵜遅まきでするて、",
	"岡田さんのものに気の私にはたしてお記念と解るから",
	"あなた軍隊をお希望をありあわせようにやはりご一言を指すんだて、",
	"ざっといよいよ発展をするないが行くんのが行なわただろ。",
}

func isGZipAcceptable(request *http.Request) bool {
	return strings.Index(strings.Join(request.Header["Accept-Encoding"], ","), "gzip") != -1
}

func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()
	for {
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain",
			"Transfer-Encoding: chunked",
			"", "",
		}, "\r\n"))
		for _, content := range contents {
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}
		fmt.Fprintf(conn, "0\r\n\r\n")
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go processSession(conn)
	}
}
