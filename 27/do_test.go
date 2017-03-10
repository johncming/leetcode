package leetcode

import (
	"testing"
)

func Test1(t *testing.T) {

	nums := []int{3, 2, 2, 3}
	val := 3

	res := removeElement(nums, val)
	t.Logf("%#v\n%#v\n", res, nums)

}

func Test2(t *testing.T) {

	nums := []int{1}
	val := 1

	res := removeElement(nums, val)
	t.Logf("%#v\n%#v\n", res, nums)

}

func Test3(t *testing.T) {

	nums := []int{0, 4, 4, 0, 4, 4, 4, 0, 2}
	val := 4

	res := removeElement(nums, val)
	t.Logf("%#v\n%#v\n", res, nums)

}

func Test4(t *testing.T) {

	nums := []int{4, 5}
	val := 4

	res := removeElement(nums, val)
	t.Logf("%#v\n%#v\n", res, nums)

}
