package arrays

func urlify(s string, l int) string {
	sarr := []rune(s)
	current := l - 1
	insertAt := len(s) - 1

	for current >= 0 {
		if sarr[current] != ' ' {
			sarr[insertAt] = sarr[current]
			insertAt--
		} else {
			sarr[insertAt] = '0'
			sarr[insertAt-1] = '2'
			sarr[insertAt-2] = '%'
			insertAt -= 3
		}

		current--
	}

	return string(sarr)
}
