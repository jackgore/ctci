package arrays

import (
	"sort"
	"strings"
)

/* isUniqueLoops determines if the given
 * string contains any duplicate characters.
 * Runtime: O(n^2)
 */
func isUniqueLoops(s string) bool {
	for i, val := range s {
		if contains(s, val, i) {
			return false
		}
	}

	return true
}

/* contains is a helper function for isUnqiueLoops
 * determins if the given string s contains character
 * c anywhere except for index j.
 */
func contains(s string, c rune, j int) bool {
	for i, val := range s {
		if val == c && i != j {
			return true
		}
	}

	return false
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// O(nlog(n))
func isUniqueSort(s string) bool {
	s = sortString(s)
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return false
		}
	}

	return true
}

func isUniqueMap(s string) bool {
	set := make(map[rune]bool)

	for _, val := range s {
		if _, ok := set[val]; ok {
			return false
		} else {
			set[val] = true
		}
	}

	return true
}
