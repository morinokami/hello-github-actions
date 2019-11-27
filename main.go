package main

import (
	"fmt"
	"io"
	"os"
)

// Lang represents the language used to greet.
type Lang int

// En: English, Zh: Chinese
const (
	En Lang = iota
	Zh
)

// Hello greets GitHub Actions.
func Hello(out io.Writer, lang Lang) {
	if lang == Zh {
		fmt.Fprint(out, "你好, GitHub Actions")
	}
	fmt.Fprint(out, "Hello, GitHub Actions!")
}

func main() {
	Hello(os.Stdout, En)
}
