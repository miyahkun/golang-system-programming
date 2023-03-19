package main

import (
	"fmt"
	"io"
	"os"
)

type PascalCaseWriter struct {
	w    io.Writer
	last byte
}

func (w *PascalCaseWriter) Write(p []byte) (int, error) {
	r := 0
	var b [1]byte
	for n := range p {
		b[0] = p[n]
		switch w.last {
		case ' ', '\t', '\r', '\n', 0:
			if 'a' <= b[0] && b[0] <= 'z' {
				b[0] -= 32
			}
		}

		nw, err := w.w.Write(b[:])
		if err != nil {
			return r + nw, err
		}
		w.last = b[0]
	}
	return r, nil
}

func NewPascalCaseWriter(w io.Writer) *PascalCaseWriter {
	return &PascalCaseWriter{w, 0}
}

func main() {
	w := NewPascalCaseWriter(os.Stdout)
	fmt.Fprintln(w, "hello world")
}
