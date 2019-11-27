package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Hello in English", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		lang := En

		Hello(buffer, lang)

		got := buffer.String()
		want := "Hello, GitHub Actions!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Chinese", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		lang := Zh

		Hello(buffer, lang)

		got := buffer.String()
		want := "你好, GitHub Actions!"

		assertCorrectMessage(t, got, want)
	})

}
