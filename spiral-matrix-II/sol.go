
func matrixTranspose(matrix [][]int) [][]int {
	newMatrix := [][]int{}
	h := len(matrix)

	if h == 0 {
		return newMatrix
	}

	w := len(matrix[0])

	for i := 0; i < w; i++ {
		row := []int{}
		for j := 0; j < h; j++ {

			row = append(row, matrix[h-j-1][i])
		}

		newMatrix = append(newMatrix, row)
	}
	return newMatrix
}

func firstZero(matrix []int) (int, bool) {

	for i := len(matrix) - 1; i >= 0; i-- {
		if matrix[i] == 0 {
			return i, true
		}
	}

	return 0, false
}

func generateMatrix(n int) [][]int {

	result := [][]int{}

	max := n * n

	for i := max; i > 0; i-- {

		if len(result) == 0 {
			result = append(result, []int{i})
		} else {

			if k, ok := firstZero(result[0]); ok {
				result[0][k] = i
			} else {
				for j := 0; j < len(result); j++ {

					if j == 0 {

						result[j] = append([]int{i}, result[j]...)

					} else {
						result[j] = append([]int{0}, result[j]...)
					}

				}
				result = matrixTranspose(result)
			}
		}

	}

	return result

}
