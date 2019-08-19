func checkDuplicated(groups [][]byte) bool {

	for _, group := range groups {
		numbers := map[byte]struct{}{}
		for _, n := range group {
			if _, ok := numbers[n]; ok {
				return true
			} else {
				numbers[n] = struct{}{}
			}

		}
	}
	return false
}

func isValidSudoku(board [][]byte) bool {
	// 49 (1) ~ 57 (9)

	rows := make([][]byte, 9)
	cols := make([][]byte, 9)
	blocks := make([][]byte, 9) // 0 ~ 9

	for y, line := range board {

		for x, num := range line {
			if int(num) == 46 {
				continue
			}

			rows[x] = append(rows[x], num)
			cols[y] = append(cols[y], num)

			_y := y / 3
			blockIndex := 3*_y + x/3
			blocks[blockIndex] = append(blocks[blockIndex], num)

		}
	}

	return !(checkDuplicated(rows) || checkDuplicated(cols) || checkDuplicated(blocks))
}

