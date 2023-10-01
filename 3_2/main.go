package main

import (
	"fmt"
	"io"
	"log"
	"os"
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

	buf2 := make([]byte, 500)
	size2, err2 := io.ReadFull(r2, buf2)

	if err2 == io.ErrUnexpectedEOF {
		fmt.Printf("%s (%d bytes read)\n", buf2, size2)
	} else {
		log.Fatal(err2)
	}

	fmt.Printf("Size: %d, Content: %s\n", size2, buf2)

	fmt.Println("----- With io.Copy -----")

	r3 := strings.NewReader("some io.Reader stream to be read\n")

	if _, err3 := io.Copy(os.Stdout, r3); err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println("----- With io.CopyN -----")

	r4 := strings.NewReader("some io.Reader stream to be read\n")

	if _, err4 := io.CopyN(os.Stdout, r4, 8); err4 != nil {
		log.Fatal(err4)
	}
}
