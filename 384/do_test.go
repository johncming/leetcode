package leetcode

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	nums := []int{1, 2, 3}

	obj := Constructor(nums)

	fmt.Printf("%+v\n", obj.Reset())   // output for debug
	fmt.Printf("%+v\n", obj.Shuffle()) // output for debug

}
