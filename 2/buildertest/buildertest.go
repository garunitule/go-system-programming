package buildertest

import (
	"fmt"
	"strings"
)

func WriteToStringsBuffer() {
	var builder strings.Builder
	builder.Write([]byte("string.Builder example\n"))
	fmt.Println(builder.String())
}
