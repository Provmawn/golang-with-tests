package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	checkSums := func(t *testing.T, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	}
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		checkSums(t, got, want, numbers)
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		checkSums(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("accurate slice containing the sum of each slice", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("accurate sum for the tail slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4}, []int{9, 8, 7})
		want := []int{9, 15}
		checkSums(t, got, want)
	})
	t.Run("passing in an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{})
		want := []int{0}
		checkSums(t, got, want)
	})
}

func TestArray(t *testing.T) {
	assertEqualInt := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
	t.Run("access nth element of initialized array", func(t *testing.T) {
		arr := [3]int{1, 2, 3}
		got := arr[2]
		want := 3
		assertEqualInt(t, got, want)
	})
	t.Run("access the nth element a zero-value initialized array", func(t *testing.T) {
		var arr [3]int
		got := arr[0]
		want := 0
		assertEqualInt(t, got, want)
	})
	t.Run("letting the compiler decide on the size of an array", func(t *testing.T) {
		arr := [...]int{1, 2, 3, 4}
		got := len(arr)
		want := 4
		assertEqualInt(t, got, want)
	})
}

func TestSlice(t *testing.T) {
	assertEqualInt := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
	t.Run("capacity is the same as length when capacity is ommited", func(t *testing.T) {
		slice := make([]int, 5)
		got := cap(slice)
		want := 5
		assertEqualInt(t, got, want)
	})
	t.Run("zero value for a slice is nil", func(t *testing.T) {
		var a []int
		got := a
		if got != nil {
			t.Errorf("not nil")
		}
	})
	t.Run("length is zero for a zero value slice", func(t *testing.T) {
		var a []int
		got := len(a)
		want := 0
		assertEqualInt(t, got, want)
	})
	t.Run("creating a slice from an array", func(t *testing.T) {
		var a [3]int
		b := a[:]
		got := cap(b)
		want := 3
		assertEqualInt(t, got, want)
	})
	t.Run("slicing an array and modifying it modifies the original array", func(t *testing.T) {
		a := [3]int{1, 2, 3}
		b := a[:]
		b[0] = 5
		got := a[0]
		want := 5
		assertEqualInt(t, got, want)
	})
	t.Run("append one slice to another", func(t *testing.T) {
		slice1 := []string{"welcome", "to"}
		slice2 := []string{"the", "showdown", "!"}
		got := append(slice1, slice2...)
		want := []string{"welcome", "to", "the", "showdown", "!"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("copy slice over", func(t *testing.T) {
		slice := make([]byte, 1e6)
		b := make([]byte, 2)
		copy(b, slice[:2])
		got := len(b)
		want := 2
		assertEqualInt(t, got, want)
	})
}
