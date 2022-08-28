package maps

import (
	"testing"
)

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("got error searching dictionary:", err)
	}
	if got != definition {
		t.Errorf("got %s, want %s for definition of %s", got, definition, word)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": testDefinition}

	t.Run("search for a word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := testDefinition
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := errNoKey
		if err == nil {
			t.Fatal("was expecting an error")
		}
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("adding a word to the dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", testDefinition)
		if err != nil {
			t.Errorf("word already in dictionary")
		}
		assertDefinition(t, dictionary, "test", testDefinition)
	})
	t.Run("word already exists in dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": testDefinition}
		err := dictionary.Add("test", testDefinition)
		assertError(t, err, errExistingElement)
		assertDefinition(t, dictionary, "test", testDefinition)
	})
}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{"test": "test is a test"}
	t.Run("update definition of a word", func(t *testing.T) {
		err := dictionary.Update("test", testDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, "test", testDefinition)
	})

	t.Run("update word that does not exist", func(t *testing.T) {
		err := dictionary.Update("hello", "used as a greeting")
		assertError(t, err, errNoKey)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete word from dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": testDefinition}
		dictionary.Delete("test")
		_, err := dictionary.Search("test")
		assertError(t, err, errNoKey)
	})
}

func TestMap(t *testing.T) {
	t.Run("declare map is nil", func(t *testing.T) {
		var m map[int]int
		if m != nil {
			t.Errorf("expected to be nil")
		}
	})
	t.Run("make map is not nil", func(t *testing.T) {
		m := make(map[int]int)
		if m == nil {
			t.Errorf("should not be nil")
		}
	})
	t.Run("zero value of map if key doesn't exist", func(t *testing.T) {
		var m map[string]string
		got := m["test"]
		want := ""
		if got != want {
			t.Errorf("expected %s, got %s", got, want)
		}
	})
	t.Run("get length of map returns the number of elements", func(t *testing.T) {
		ages := make(map[string]int)
		ages["master chief"] = 41
		ages["kratos"] = 1050
		got := len(ages)
		want := 2
		if got != want {
			t.Errorf("expected length of map to be %d", want)
		}
	})
	t.Run("delete will remove an element from map", func(t *testing.T) {
		ages := make(map[string]int)
		ages["master chief"] = 41
		ages["kratos"] = 1050
		delete(ages, "master chief")
		got := len(ages)
		want := 1
		if got != want {
			t.Errorf("expected length of map to be %d", want)
		}
	})
	t.Run("check if a key exists", func(t *testing.T) {
		ages := make(map[string]int)
		ages["master chief"] = 41
		ages["kratos"] = 1050
		_, ok := ages["kratos"]
		if ok != true {
			t.Errorf("key should exist")
		}
	})
}
