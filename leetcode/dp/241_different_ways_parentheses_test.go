package dp

import (
	"fmt"
	"github.com/sokrato/goal/leetcode/util"
	"strings"
)

func splitExpr(expr string) (res []string) {
	sb := strings.Builder{}
	for _, ch := range expr {
		if '0' <= ch && ch <= '9' {
			sb.WriteRune(ch)
		} else {
			res = append(res, sb.String())
			sb.Reset()
			res = append(res, string(ch))
		}
	}
	res = append(res, sb.String())
	return
}

var ops = map[string]func(a, b int) int{
	"+": func(a, b int) int {
		return a + b
	},
	"-": func(a, b int) int {
		return a - b
	},
	"*": func(a, b int) int {
		return a * b
	},
}

// 241. Different Ways to Add Parentheses
// https://leetcode.com/problems/different-ways-to-add-parentheses/description/
func diffWaysToCompute(expression string) (res []int) {
	exprs := splitExpr(expression)
	mem := map[string][]int{}

	prev := util.ParseInt(exprs[0])
	mem[exprs[0]] = []int{prev}
	maxNumOps := len(exprs) >> 1
	for numOps := 0; numOps < maxNumOps; numOps++ {
		//
	}
	for i := 1; i < len(exprs); i += 2 {
		next := util.ParseInt(exprs[i+1])
		mem[exprs[i+1]] = []int{next}
		term := fmt.Sprintf("%v%v%v", prev, exprs[i], next)
		val := ops[exprs[i]](prev, next)
		mem[term] = []int{val}

		prev = next
	}
	return mem[expression]
}
