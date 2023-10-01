package main

import (
	"io"
	"os"
	"path/filepath"
)

func main() {
	path, err1 := os.Executable()
	if err1 != nil {
		panic(err1)
	}
	dir := filepath.Dir(path)
	file, err2 := os.Open(dir + "/" + "file.go")
	if err2 != nil {
		panic(err2)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
