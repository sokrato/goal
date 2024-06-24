package array

import "testing"

// 59. Spiral Matrix II
// https://leetcode.cn/problems/spiral-matrix-ii/description/
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for idx := range res {
		res[idx] = make([]int, n)
	}

	end := n*n + 1
	di := 0
	x, y := 0, 0
	for i := 1; i < end; i++ {
		res[x][y] = i

		dir := dirs[di%4]
		x2 := x + dir[0]
		y2 := y + dir[1]
		if x2 < 0 || y2 < 0 || x2 >= n || y2 >= n || res[x2][y2] > 0 {
			di++
			dir = dirs[di%4]
			x2 = x + dir[0]
			y2 = y + dir[1]
		}
		x, y = x2, y2
	}
	return res
}

func Test_generateMatrix(t *testing.T) {
	for i := 1; i < 3; i++ {
		generateMatrix(i)
	}
}
