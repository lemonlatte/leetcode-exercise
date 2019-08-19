
func matrixTranspose(matrix [][]int) [][]int {
	newMatrix := [][]int{}
	h := len(matrix)

	if h == 0 {
		return newMatrix
	}

	w := len(matrix[0])

	for i := w - 1; i >= 0; i-- {
		row := []int{}
		for j := 0; j < h; j++ {

			row = append(row, matrix[j][i])
		}

		newMatrix = append(newMatrix, row)
	}
	return newMatrix
}

func spiralOrder(matrix [][]int) []int {
	list := []int{}

	for len(matrix) != 0 {

		list = append(list, matrix[0]...)
		matrix = matrixTranspose(matrix[1:])
	}
	return list
}
