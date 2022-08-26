package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(5, 5)
	expected := 10
	assertEqual(t, got, expected)
}

func ExampleAdd() {
	sum := Add(2, 2)
	fmt.Println(sum)
	// output: 4
}

func assertEqual(t testing.TB, got int, expected int) {
	t.Helper()
	if got != expected {

		t.Errorf("got %d, expected %d", got, expected)
	}
}
