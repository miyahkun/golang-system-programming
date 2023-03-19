package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "World",
	}

	jsonBytes, err := json.Marshal(source)
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
		panic(err)
	}

	gzipWriter := gzip.NewWriter(w)
	multiwriter := io.MultiWriter(gzipWriter, os.Stdout)
	io.WriteString(multiwriter, string(jsonBytes))
	gzipWriter.Close()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
