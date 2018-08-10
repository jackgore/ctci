package arrays

import (
	"fmt"
	"strings"
)

func compression(s string) string {
	c := compress(s)

	if len(c) >= len(s) {
		return s
	}

	return c
}

func compress(s string) string {
	var b strings.Builder
	i := 0
	for i < len(s) {
		cur := s[i]
		count := 0
		for i < len(s) && s[i] == cur {
			count++
			i++
		}

		fmt.Fprintf(&b, "%v%v", string(cur), count)
	}

	return b.String()
}
