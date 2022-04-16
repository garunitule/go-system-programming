package textread

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Read() {
	var source = `1行目
	2行目
	3行目`
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}
}
