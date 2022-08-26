package iteration

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

func TestRepeatOld(t *testing.T) {
	got := Repeat('a', 5)
	want := "aaaaa"
	if got != want {
		t.Errorf("got %q, expected %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat('a', 5)
	}
}

//hello
func ExampleRepeat() {
	linebreak := Repeat('=', 5)
	fmt.Println(linebreak)
}

func TestCompare(t *testing.T) {
	t.Run("Less Than", func(t *testing.T) {
		got := strings.Compare("Amman", "Mason")
		want := -1
		assertEqualInt(got, want, t)
	})
	t.Run("Greater Than", func(t *testing.T) {
		got := strings.Compare("Tyler", "Emmett")
		want := 1
		assertEqualInt(got, want, t)
	})
	t.Run("Equal", func(t *testing.T) {
		got := strings.Compare("Amman", "Amman")
		want := 0
		assertEqualInt(got, want, t)
	})
}

func TestContains(t *testing.T) {
	t.Run("contains", func(t *testing.T) {
		got := strings.Contains("Goodbye", "Goo")
		want := true
		assertEqualBool(got, want, t)
	})
	t.Run("does not contain", func(t *testing.T) {
		got := strings.Contains("Nope", "Not here")
		want := false
		assertEqualBool(got, want, t)
	})
}

func TestContainsAny(t *testing.T) {
	t.Run("does contain", func(t *testing.T) {
		got := strings.ContainsAny("hello", "ei")
		want := true
		assertEqualBool(got, want, t)
	})
	t.Run("does not contain", func(t *testing.T) {
		got := strings.ContainsAny("hello", "wat")
		want := false
		assertEqualBool(got, want, t)
	})
}

func TestContainsRune(t *testing.T) {
	t.Run("does contain", func(t *testing.T) {
		got := strings.ContainsRune("ammn", 97)
		want := true
		assertEqualBool(got, want, t)
	})
	t.Run("does not contain", func(t *testing.T) {
		got := strings.ContainsRune("yo", 97)
		want := false
		assertEqualBool(got, want, t)
	})
}

func TestCount(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		got := strings.Count("hello", "")
		want := len("hello") + 1
		assertEqualInt(got, want, t)
	})
	t.Run("non-empty", func(t *testing.T) {
		got := strings.Count("I want to the store and bought the pizza.", "the")
		want := 2
		assertEqualInt(got, want, t)
	})
}

func TestCut(t *testing.T) {
	t.Run("correct before", func(t *testing.T) {
		got, _, _ := strings.Cut("the large cat", "large")
		want := "the "
		assertEqualString(got, want, t)
	})
	t.Run("correct after", func(t *testing.T) {
		_, got, _ := strings.Cut("the large cat", "large")
		want := " cat"
		assertEqualString(got, want, t)
	})
	t.Run("found", func(t *testing.T) {
		_, _, got := strings.Cut("the large cat", "large")
		want := true
		assertEqualBool(got, want, t)
	})
	t.Run("not found", func(t *testing.T) {
		_, _, got := strings.Cut("the large cat", "foo")
		want := false
		assertEqualBool(got, want, t)
	})
}

func TestEqualFold(t *testing.T) {
	t.Run("equal, different casing", func(t *testing.T) {
		got := strings.EqualFold("the", "ThE")
		want := true
		assertEqualBool(got, want, t)
	})
	t.Run("equal, same casing", func(t *testing.T) {
		got := strings.EqualFold("the", "the")
		want := true
		assertEqualBool(got, want, t)
	})
	t.Run("not equal, different casing", func(t *testing.T) {
		got := strings.EqualFold("the", "waT")
		want := false
		assertEqualBool(got, want, t)
	})
	t.Run("not equal, same casing", func(t *testing.T) {
		got := strings.EqualFold("the", "wat")
		want := false
		assertEqualBool(got, want, t)
	})
}

func TestFields(t *testing.T) {
	t.Run("check for expected slice", func(t *testing.T) {
		got := strings.Fields("hello world")
		want := []string{"hello", "world"}
		assertEqualStringSlice(got, want, t)

	})
	t.Run("check for alternative spacing", func(t *testing.T) {
		got := strings.Fields("hello\tto the\nworld")
		want := []string{"hello", "to", "the", "world"}
		assertEqualStringSlice(got, want, t)
	})
}

func TestFieldFunc(t *testing.T) {
	t.Run("check for expected slice", func(t *testing.T) {
		f := func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}
		got := strings.FieldsFunc("--  hello, world ?", f)
		want := []string{"hello", "world"}
		assertEqualStringSlice(got, want, t)
	})
}

func TestHasPrefix(t *testing.T) {
	t.Run("check for wrong prefix", func(t *testing.T) {
		got := strings.HasPrefix("Golang", "hello")
		want := false
		assertEqualBool(got, want, t)
	})
	t.Run("check for lowercase prefix", func(t *testing.T) {
		got := strings.HasPrefix("Golang", "go")
		want := false
		assertEqualBool(got, want, t)
	})
	t.Run("check for correct prefix", func(t *testing.T) {
		got := strings.HasPrefix("Golang", "Go")
		want := true
		assertEqualBool(got, want, t)
	})
}

func TestHasSuffix(t *testing.T) {
	t.Run("check for correct suffix", func(t *testing.T) {
		got := strings.HasSuffix("script.sh", "sh")
		want := true
		assertEqualBool(got, want, t)
	})
	t.Run("check for wrong suffix", func(t *testing.T) {
		got := strings.HasSuffix("script.sh", "txt")
		want := false
		assertEqualBool(got, want, t)
	})
}

func TestIndex(t *testing.T) {
	t.Run("check found", func(t *testing.T) {
		got := strings.Index("hello world", " ")
		want := 5
		assertEqualInt(got, want, t)
	})
	t.Run("check not found", func(t *testing.T) {
		got := strings.Index("hello world", "foo")
		want := -1
		assertEqualInt(got, want, t)
	})
}

func TestIndexAny(t *testing.T) {
	t.Run("check has unicode code point", func(t *testing.T) {
		got := strings.IndexAny("hello world", "aieou")
		want := 1
		assertEqualInt(got, want, t)
	})
	t.Run("check does not have unicode code point", func(t *testing.T) {
		got := strings.IndexAny("hello world", "a")
		want := -1
		assertEqualInt(got, want, t)
	})
}

func TestIndexByte(t *testing.T) {
	t.Run("check ascii character found ", func(t *testing.T) {
		got := strings.IndexByte("hello", 'e')
		want := 1
		assertEqualInt(got, want, t)
	})
	t.Run("check ascii character not found", func(t *testing.T) {
		got := strings.IndexByte("hello", 'f')
		want := -1
		assertEqualInt(got, want, t)
	})
}

func TestJoin(t *testing.T) {
	t.Run("check join is correct", func(t *testing.T) {
		got := strings.Join([]string{"Hello", "world"}, " - ")
		want := "Hello - world"
		assertEqualString(got, want, t)
	})
}

func TestLastIndex(t *testing.T) {
	t.Run("test found index", func(t *testing.T) {
		got := strings.LastIndex("world world", "world")
		want := 6
		assertEqualInt(got, want, t)
	})
	t.Run("could not find", func(t *testing.T) {
		got := strings.LastIndex("hello world", "the")
		want := -1
		assertEqualInt(got, want, t)
	})
}

func TestMap(t *testing.T) {
	t.Run("check if character swapping works", func(t *testing.T) {
		mapping := func(r rune) rune {
			switch {
			case r >= 'A' && r <= 'X':
				return r + 2
			case r >= 'a' && r <= 'x':
				return r + 2
			}
			return r
		}
		got := strings.Map(mapping, "hello world")
		want := "jgnnq yqtnf"
		assertEqualString(got, want, t)
	})
	t.Run("check drop character", func(t *testing.T) {
		mapping := func(r rune) rune {
			if r != 'e' && r != 'w' {
				return -1
			}
			return r
		}
		got := strings.Map(mapping, "hello world")
		want := "ew"
		assertEqualString(got, want, t)
	})
}

func TestRepeat(t *testing.T) {
	t.Run("check repeat 5 times", func(t *testing.T) {
		got := strings.Repeat("hello", 5)
		want := "hellohellohellohellohello"
		assertEqualString(got, want, t)
	})
	t.Run("check panic on receiving -1", func(t *testing.T) {
		assertPanic(strings.Repeat, "hello", -1, t)
	})
}

func assertEqualBool(got, want bool, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func assertEqualInt(got, want int, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertEqualString(got, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertEqualStringSlice(got, want []string, t *testing.T) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("slices do not have the same size")
	}
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("got arr[%d] = %q, want arr[%d] = %q", i, got[i], i, want[i])
		}
	}
}

func assertPanic(f func(string, int) string, s string, i int, t *testing.T) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f(s, i)
}
