package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("----- stdin -----")

	for {
		buffer := make([]byte, 5)
		size1, err1 := os.Stdin.Read(buffer)
		if err1 == io.EOF {
			fmt.Println("\nEOF")
			break
		}
		fmt.Printf("size=%d input='%s'\n", size1, string(buffer))
	}
}
