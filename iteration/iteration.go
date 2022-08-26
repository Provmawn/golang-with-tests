package iteration

import (
	"strings"
)

func Repeat(c rune, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += string(c)
	}
	return result
}

func SameNames(name1, name2 string) bool {
	if strings.Compare(name1, name2) != 0 {
		return false
	}
	return true
}
