package main

import (
	"fmt"
	"io"
	"os"
)

// Hello greets GitHub Actions.
func Hello(out io.Writer) {
	fmt.Fprint(out, "Hello, GitHub Actions!")
}

func main() {
	Hello(os.Stdout)
}
