package arrays

func zero(a [][]int) [][]int {
	n := len(a)
	columns := make(map[int]bool)
	rows := make(map[int]bool)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] == 0 {
				columns[i] = true
				rows[j] = true
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			_, ok1 := columns[i]
			_, ok2 := rows[j]
			if ok1 || ok2 {
				a[i][j] = 0
			}
		}
	}

	return a
}
