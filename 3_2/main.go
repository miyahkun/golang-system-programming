package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	fmt.Println("With io.ReadAll")
	r := strings.NewReader("Go is general-purpose language designed with systems programming in mind.")

	buffer, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", buffer)
}
