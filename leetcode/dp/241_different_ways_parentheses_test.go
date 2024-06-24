package dp

import (
	"fmt"
	"testing"
	"unicode"
)

func splitExpr(expr string) (numbers []int, operators []rune) {
	val := 0
	for _, ch := range expr {
		if unicode.IsDigit(ch) {
			val = val*10 + int(ch-'0')
		} else {
			numbers = append(numbers, val)
			val = 0
			operators = append(operators, ch)
		}
	}
	numbers = append(numbers, val)
	return
}

// 241. Different Ways to Add Parentheses
// https://leetcode.com/problems/different-ways-to-add-parentheses
// TODO: 如何用迭代形式去实现？
// func diffWaysToCompute(expression string) []int {
// 	exprs := splitExpr(expression)
// 	mem := map[string][]int{}

// 	prev := util.ParseInt(exprs[0])
// 	mem[exprs[0]] = []int{prev}
// 	for i := 1; i < len(exprs); i += 2 {
// 		next := util.ParseInt(exprs[i+1])
// 		mem[exprs[i+1]] = []int{next}

// 		term := fmt.Sprintf("%v%v%v", prev, exprs[i], next)
// 		val := ops[exprs[i]](prev, next)
// 		mem[term] = []int{val}

// 		prev = next
// 	}

// 	for size := 3; size < len(exprs); size += 2 {
// 		// 1+2*3-4
// 		// sliding window of 3:
// 		//   1. (1+2)*3-4
// 		//   2. 1+(2*3)-4
// 		//   3. 1+2*(3-4)
// 		// sliding window of 5:
// 		//   1. (1+2*3)-4
// 		//   2. 1+(2*3-4)
// 		half := (size - 1) >> 1
// 		for i := half; i+half <= len(exprs); i += 2 {
// 			win := strings.Join(exprs[i-half:i+half+1], "")
// 			winVals := mem[win]

// 			prevOpIndex := i - half - 1
// 			if prevOpIndex > 0 {
// 				prevOp := exprs[prevOpIndex]
// 				prevIt := exprs[prevOpIndex-1]
// 				prev := mem[prevIt][0]
// 				fn := ops[prevOp]
// 				win2 := fmt.Sprintf("%v%v%v", exprs[prevOpIndex-1], exprs[prevOpIndex], win)
// 				for _, val := range winVals {
// 					mem[win2] = append(mem[win2], fn(prev, val))
// 				}
// 			}

// 			nextOpIndex := i + half + 1
// 			if nextOpIndex+1 < len(exprs) {
// 				nextOp := exprs[nextOpIndex]
// 				nextIt := exprs[nextOpIndex+1]
// 				next := mem[nextIt][0]
// 				fn := ops[nextOp]
// 				win2 := fmt.Sprintf("%v%v%v", win, nextOp, nextIt)
// 				for _, val := range winVals {
// 					mem[win2] = append(mem[win2], fn(val, next))
// 				}
// 			}
// 		}
// 	}

// 	// 组合
// 	res := mem[expression]
// 	for i := 3; i < len(exprs)-2; i += 2 {
// 		prevIt := strings.Join(exprs[0:i], "")
// 		nextIt := strings.Join(exprs[i+1:], "")
// 		fn := ops[exprs[i]]
// 		for _, prev := range mem[prevIt] {
// 			for _, next := range mem[nextIt] {
// 				res = append(res, fn(prev, next))
// 			}
// 		}
// 	}
// 	return res
// }

func diffWaysToComputeRecur(expression string) []int {
	numbers, operators := splitExpr(expression)
	mem := map[int]map[int][]int{} // lo - hi - []int
	for i := 0; i < len(numbers); i++ {
		mem[i] = map[int][]int{}
	}

	var dfs func(lo, hi int) []int
	dfs = func(lo, hi int) []int { // hi - exclusive
		if lo+1 == hi {
			val := numbers[lo]
			return []int{val}
		}

		if res := mem[lo][hi]; res != nil {
			return res
		}

		var res []int
		for i := lo; i < hi-1; i++ {
			left := dfs(lo, i+1)
			right := dfs(i+1, hi)
			for _, x := range left {
				for _, y := range right {
					tmp := 0
					switch operators[i] {
					case '+':
						tmp = x + y
					case '-':
						tmp = x - y
					default:
						tmp = x * y
					}
					res = append(res, tmp)
				}
			}
		}
		mem[lo][hi] = res
		return res
	}
	return dfs(0, len(numbers))
}

func Test_diffWaysToCompute(t *testing.T) {
	vals := diffWaysToComputeRecur("2*3-4*5")
	for _, val := range vals {
		fmt.Println(val)
	}
}
