package pascal

func valOr0(arr [][]int, i int, j int) int {
	if i >= 0 && i < len(arr) && j >= 0 && j < len(arr[i]) {
		return arr[i][j]
	}

	return 0
}

// Triangle computes the pascal triangle to the Nth row
// and returns the result in a 2d array
func Triangle(n int) [][]int {
	var res [][]int

	for i := 0; i < n; i++ {
		var row []int
		row = append(row, 1)
		for j := 1; j <= i; j++ {
			row = append(row, valOr0(res, i-1, j-1) + valOr0(res, i-1, j))
		}
		res = append(res, row)
	}
	return res
}
