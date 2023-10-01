package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	fmt.Println("----- With io.ReadAll -----")
	r1 := strings.NewReader("Go is general-purpose language designed with systems programming in mind.")

	buf1, err1 := io.ReadAll(r1)
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Printf("%s\n", buf1)

	fmt.Println("----- With io.ReadFull -----")

	r2 := strings.NewReader("Go is general-purpose language designed with systems programming in mind.")

	buf2 := make([]byte, 5)
	size2, err2 := io.ReadFull(r2, buf2)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Printf("Size: %d, Content: %s\n", size2, buf2)
}
