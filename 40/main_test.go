package leetcode

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	res := combinationSum2(candidates, 8)
	fmt.Printf("%+v\n", res) // output for debug
}
