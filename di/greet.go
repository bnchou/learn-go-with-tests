package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(w io.Writer, s string) {
	fmt.Fprintf(w, "Hello, %s", s)
}

func main() {
	Greet(os.Stdout, "Foo\n")
}
