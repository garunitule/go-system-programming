package bufferread

import (
	"bytes"
	"io"
	"os"
)

func Read() {
	// var buffer1 bytes.Buffer
	// buffer2 := bytes.NewBuffer([]byte{0x10, 0x20, 0x30})
	buffer3 := bytes.NewBufferString("初期文字列")
	io.Copy(os.Stdout, buffer3)
}
