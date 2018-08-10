package arrays

import (
	"strings"
)

func createLetterCount(s string) map[rune]int {
	m := make(map[rune]int)

	for _, val := range s {
		lower := strings.ToLower(string(val))

		if val, ok := m[[]rune(lower)[0]]; !ok {
			m[[]rune(lower)[0]] = 1
		} else {
			m[[]rune(lower)[0]] = val + 1
		}
	}

	return m
}

func palindrome(s string) bool {
	m := createLetterCount(s)
	foundOdd := false

	for key, value := range m {
		if key == ' ' {
			continue
		}

		if (value % 2) == 1 {
			if foundOdd {
				return false
			}

			foundOdd = true
		}
	}

	return true
}
