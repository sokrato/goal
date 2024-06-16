package unionfind

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 886. Possible Bipartition
// https://leetcode.com/problems/possible-bipartition/
// 1. build the graph
// 2. color the nodes without conflicts
func possibleBipartition(n int, dislikes [][]int) bool {
	const (
		unset = 0
		black = 1
		//white = 2
		all = 3
	)
	colors := make([]int, n)
	conns := make([][]int, n)

	for _, pair := range dislikes {
		a, b := pair[0], pair[1]
		conns[a-1] = append(conns[a-1], b)
		conns[b-1] = append(conns[b-1], a)
	}

	// 必须递归上色，否则因为顺序原因，上色不准
	var dfs func(person, lastColor int) bool
	dfs = func(person, lastColor int) bool {
		i := person - 1
		thisColor := colors[i]
		if thisColor != unset { // 已经染色成功，只需要判断合法性
			return thisColor != lastColor
		}

		// 染色
		if lastColor == unset {
			thisColor = black
		} else {
			thisColor = all - lastColor
		}
		colors[i] = thisColor

		// 递归下探
		for _, next := range conns[i] {
			if !dfs(next, thisColor) {
				return false
			}
		}
		return true
	}
	for x := range conns {
		if !dfs(x+1, unset) {
			return false
		}
	}
	return true
}

func Test_possibleBipartition(t *testing.T) {
	cases := []struct {
		n        int
		dislikes [][]int
		expected bool
	}{
		{
			n:        3,
			dislikes: [][]int{{1, 2}, {1, 3}, {2, 3}},
			expected: false,
		},
		{
			n:        4,
			dislikes: [][]int{{1, 2}, {1, 3}, {2, 4}},
			expected: true,
		},
		{
			n:        5,
			dislikes: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}},
			expected: false,
		},
		{
			n:        10,
			dislikes: [][]int{{4, 7}, {4, 8}, {5, 6}, {1, 6}, {3, 7}, {2, 5}, {5, 8}, {1, 2}, {4, 9}, {6, 1}, {8, 1}, {3, 6}, {2, 1}, {9, 1}, {3, 9}, {2, 3}, {1, 9}, {4, 6}, {5, 7}, {3, 8}, {1, 8}, {1, 7}, {2, 4}},
			expected: true,
		},
	}

	a := assert.New(t)
	for i, it := range cases {
		a.Equal(it.expected, possibleBipartition(it.n, it.dislikes), i)
	}
}
