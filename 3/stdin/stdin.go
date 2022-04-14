package stdin

import (
	"fmt"
	"io"
	"os"
)

func Read() {
	for {
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF\n")
			break
		}

		fmt.Printf("size=%d input='%s'\n", size, string(buffer))
	}
}
