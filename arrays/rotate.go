package arrays

// Swap columns with rows
func transpose(a [][]int) [][]int {
	offset := 0
	n := len(a)

	for i := 0; i < n; i++ {
		for j := offset; j < n; j++ {
			a[i][j], a[j][i] = a[j][i], a[i][j]
		}

		offset++
	}

	return a
}

func verticalFlip(a [][]int) [][]int {
	n := len(a)
	for j := 0; j < n; j++ {
		for i := 0; i < n/2; i++ {
			flip := n - (i + 1)
			a[i][j], a[flip][j] = a[flip][j], a[i][j]
		}
	}

	return a
}

func rotate(a [][]int) [][]int {
	return verticalFlip(transpose(a))
}

/*func rotateSwap(a [][]int) [][]int {
	n := len(a)
	layers := (n + 1) / 2

	for i := 0; i < layers; i++ {
		for j := 0; j < n-(i+1); j++ {
			a[i][j+i], a[j+i][i], a[n-(i+1)][j+i], a[n-(i+1+j)][n-(i+1)] = a[n-(i+1+j)][n-(i+1)], a[i][j+i], a[j+i][i], a[n-(i+1)][j+i]
		}
	}

	return a
}*/
/*
func rotateSwap(a [][]int) [][]int {
	n := len(a)
	layers := n / 2

	for i := 0; i < layers; i++ {
		first := i
		last := n - i - 1
		for j := first; j < last; j++ {
			top := a[j][first] // First column is x coord so it always starts at 'first'

			left := a[first][n-j-1]
			a[j][first] = left

			bottom := a[n-j-1][first]
			a[first][n-j-1] = bottom

			right := a[first][j]
			a[n-j-1][first] = right

			a[first][j] = top

		}
	}

	return a
}*/

func rotateSwap(a [][]int) [][]int {
	n := len(a)

	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first

			top := a[i][first] // First column is x coord so it always starts at 'first'

			// Left -> top
			a[i][first] = a[first][last-offset]

			// bottom -> left
			a[first][last-offset] = a[last-offset][last]

			// right -> bottom
			a[last-offset][last] = a[last][i]

			// top -> bottom
			a[last][i] = top

		}
	}

	return a
}
