package geometry

import (
	"fmt"
	"testing"
)

// 149. Max Points on a Line
// https://leetcode.cn/problems/max-points-on-a-line/
func maxPoints(points [][]int) int {
	panic("TODO")
}

type Point struct {
	x, y int
}

func TestName(t *testing.T) {
	a := Point{1, 2}
	b := a
	b.x = 9

	fmt.Println(a.x, b.x)
}
