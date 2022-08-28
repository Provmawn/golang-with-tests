package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type Operation string

const (
	write = Operation("write")
	sleep = Operation("sleep")
)

type SpyCountdownOperations struct {
	Calls []Operation
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
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
	t.Run("sleep after every print", func(t *testing.T) {
		spy := &SpyCountdownOperations{}
		Countdown(spy, spy)
		got := spy.Calls
		want := []Operation{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

type SpyTime struct {
	duration time.Duration
}

func (s *SpyTime) Sleep(d time.Duration) {
	s.duration = d
}

func (s *SpyTime) Write(b []byte) (n int, err error) {
	return
}

func TestConfigurableSleeper(t *testing.T) {
	t.Run("test 5 seconds sleep time", func(t *testing.T) {
		sleepTime := 5 * time.Second
		spyTime := &SpyTime{}
		sleeper := &ConfigurableSleeper{
			duration: sleepTime,
			sleep:    spyTime.Sleep,
		}
		sleeper.Sleep()
		got := spyTime.duration
		want := sleepTime
		if spyTime.duration != sleepTime {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
