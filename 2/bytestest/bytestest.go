package bytestest

import (
	"bytes"
	"fmt"
)

func WriteToBuffer() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example1\n"))
	buffer.Write([]byte("bytes.Buffer example2\n"))
	buffer.WriteString("bytes.Buffer example as String\n")
	fmt.Println(buffer.String())
}
