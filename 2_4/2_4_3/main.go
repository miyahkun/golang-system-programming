package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var buffer bytes.Buffer
	io.WriteString(&buffer, "bytes.Buffer example\n")
	fmt.Println(buffer.String())
}
