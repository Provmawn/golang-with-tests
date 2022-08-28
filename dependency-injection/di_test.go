package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("greeting message is correct", func(t *testing.T) {
		var buffer bytes.Buffer
		Greet(&buffer, "World")
		got := buffer.String()
		want := "Hello World"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
