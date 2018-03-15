package leetcode

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	nums := []int{1, 2, 3}
	res := subsets(nums)
	fmt.Printf("%+v\n", res) // output for debug
}
