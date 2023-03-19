package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "digit: %d, sting: %s, float: %f\n", 123, "Hello world", 56.7)
}
