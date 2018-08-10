package arrays

import "math"

// Determines if the given strings differ by at most one character
// Assumes len(s1) == len(s2)
func isReplacement(s1, s2 string) bool {
	difference := 0
	for i, val := range s1 {
		if val != rune(s2[i]) {
			difference++
		}
	}

	return difference <= 1
}

func oneaway(s1, s2 string) bool {
	if math.Abs(float64(len(s1)-len(s2))) > 1 {
		return false
	}

	if len(s1) == len(s2) {
		return isReplacement(s1, s2)
	}

	smaller, larger := s1, s2
	if len(s1) > len(s2) {
		larger = s1
		smaller = s2
	}

	mismatches := 0
	i := 0
	j := 0
	for j < len(larger) && i < len(smaller) {
		if mismatches > 1 {
			return false
		}

		if larger[j] != smaller[i] {
			mismatches++
		} else {
			i++
		}
		j++
	}

	return mismatches <= 1
}
