package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseIntSlice(t *testing.T) {
	a := assert.New(t)

	expected := []int{1, 2, 3}
	a.Equal(expected, ParseIntSlice("1,2, 3\n"))
	a.Equal(expected, ParseIntSlice("[ 1, 2, 3]"))
}
