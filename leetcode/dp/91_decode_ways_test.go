package dp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 91. Decode Ways
// https://leetcode.cn/problems/decode-ways/
func numDecodings(s string) int {
	memo := map[string]int{}
	var helper func(s string) int
	helper = func(s string) int {
		if s == "" || s[0] == '0' {
			return 0
		}

		if len(s) == 1 {
			return 1
		}

		v, ok := memo[s]
		if ok {
			return v
		}

		v = helper(s[1:])
		if s[0] == '1' || (s[0] == '2' && s[1] < '7') {
			if len(s) == 2 {
				v += 1
			} else {
				v += helper(s[2:])
			}
		}

		memo[s] = v
		return v
	}
	return helper(s)
}

func numDecodingsIter(s string) int {
	if s == "" || s[0] == '0' {
		return 0
	}

	x, y := 1, 1

	for i := 1; i < len(s); i++ {
		ch := s[i]
		z := 0
		if ch != '0' {
			z += y
		}
		if s[i-1] == '1' || (s[i-1] == '2' && ch <= '6') {
			z += x
		}
		x = y
		y = z
	}
	return y
}

func numDecodingsIter2(s string) int {
	size := len(s)
	dp := make([]int, size+1)
	dp[0] = 1
	for i := 1; i < size+1; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && (s[i-2]-'0')*10+(s[i-1]-'0') <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[size]
}

func Test_numDecodings(t *testing.T) {
	a := assert.New(t)
	cases := []struct {
		s string
		x int
	}{
		{"06", 0},
		{"12", 2},
		{"123", 3},
		{"133", 2},
	}
	for _, it := range cases {
		a.Equal(it.x, numDecodingsIter(it.s), it.s)
	}
}
