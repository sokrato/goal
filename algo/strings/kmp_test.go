package strings

import (
	"github.com/sokrato/goal/util/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func matches(s string, pattern string) int {
	if len(s) < len(pattern) {
		return -1
	}

	phash := 0
	shash := 0
	const base = 101 // bigger than s[i]
	mod := 1
	for i, ch := range pattern {
		phash = phash*base + int(ch)
		shash = shash*base + int(s[i])
		if i < len(pattern)-1 {
			mod *= base
		}
	}
	if phash == shash {
		return 0
	}

	for i := len(pattern); i < len(s); i++ {
		shash = (shash%mod)*base + int(s[i])
		log.GetLogger().Printf("phash=%v, shash=%v, i=%v", phash, shash, i)
		if shash == phash {
			return i - len(pattern) + 1
		}
	}

	return -1
}

func TestMatching(t *testing.T) {
	a := assert.New(t)
	a.Equal(0, matches("axe", "axe"))
	a.Equal(1, matches("aaxe", "axe"))
	a.Equal(4, matches("axbeaxe", "axe"))
}
