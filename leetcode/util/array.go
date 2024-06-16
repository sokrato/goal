package util

import "strings"

func ParseIntSlice(txt string) (res []int) {
	txt = strings.TrimSpace(txt)
	txt = strings.Trim(txt, "[]")

	parts := strings.Split(txt, ",")
	res = make([]int, len(parts))
	for idx, part := range parts {
		part = strings.TrimSpace(part)
		res[idx] = ParseInt(part)
	}
	return
}

func EqualSlices[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, v := range a {
		if v != b[idx] {
			return false
		}
	}
	return true
}
