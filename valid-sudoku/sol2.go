func isValidSudoku(board [][]byte) bool {
	// 49 (1) ~ 57 (9)

	rows := make([]map[byte]struct{}, 9)
	cols := make([]map[byte]struct{}, 9)
	blocks := make([]map[byte]struct{}, 9) // 0 ~ 9

	for y, line := range board {

		for x, num := range line {
			if int(num) == 46 {
				continue
			}

			if rows[x] == nil {
				rows[x] = map[byte]struct{}{}
			}
			if cols[y] == nil {
				cols[y] = map[byte]struct{}{}
			}

			if _, ok := rows[x][num]; ok {
				return false
			} else {
				rows[x][num] = struct{}{}
			}

			if _, ok := cols[y][num]; ok {
				return false
			} else {
				cols[y][num] = struct{}{}
			}

			_y := y / 3
			blockIndex := 3*_y + x/3

			if blocks[blockIndex] == nil {
				blocks[blockIndex] = map[byte]struct{}{}
			}

			if _, ok := blocks[blockIndex][num]; ok {
				return false
			} else {
				blocks[blockIndex][num] = struct{}{}
			}

		}
	}

	return true
}
