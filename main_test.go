package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	buffer := &bytes.Buffer{}
	Hello(buffer)

	got := buffer.String()
	want := "Hello, GitHub Actions!"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
