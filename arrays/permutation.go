package arrays

// Determines if the two strings are permutations of one another
// in O(nlog(n)) time. Using O(n) space as two new strings are allocated
// where n = max(len(s1), len(s2))
func permutationSort(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1 = sortString(s1)
	s2 = sortString(s2)

	/*
		NOTE: Don't need to loop, theyre just strings!
		for i, val := range s1 {
			if val != rune(s2[i]) {
				return false
			}
		}
	*/

	return s1 == s2
}

func createCharMap(s string) map[rune]int {
	m := make(map[rune]int)

	for _, key := range s {
		if val, ok := m[key]; !ok {
			m[key] = 1
		} else {
			m[key] = val + 1
		}
	}

	return m
}

// Determines if the two strings are permutations of one another
// in O(n) time. Using O(n) space to create a hashmap
func permutationMap(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m1 := createCharMap(s1)
	m2 := createCharMap(s2)

	for key, val1 := range m1 {
		if val2, ok := m2[key]; !ok || val1 != val2 {
			return false
		}
	}

	return true
}
