package leetcode

import (
	"testing"
)

func Test1(t *testing.T) {
	s := "aab"

	res := partition(s)

	t.Logf("%#v", res)
}
