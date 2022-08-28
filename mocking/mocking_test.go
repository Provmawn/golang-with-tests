package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	t.Run("print 3", func(t *testing.T) {
		var buffer bytes.Buffer
		sleeper := &SpySleeper{}
		Countdown(&buffer, sleeper)
		got := buffer.String()
		want := `3
2
1
Go!
`
		if got != want {
			t.Errorf("got %q, expected %q", got, want)
		}
		new_got := sleeper.Calls
		new_want := 3
		if got != want {
			t.Errorf("got %q, expected %q", new_got, new_want)
		}
	})
}
