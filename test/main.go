package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"

	jsonx86 "github.com/HobbyOSs/json-x86-64-go-mod"
)

func main() {
	r, err := gzip.NewReader(bytes.NewReader(jsonx86.X86JSONGZ))
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	contents, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Decoded JSON (%d bytes):\n%s\n", len(contents), contents[:200])
}
