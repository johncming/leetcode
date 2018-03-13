package leetcode

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	res := combinationSum(candidates, 7)
	fmt.Printf("%+v\n", res) // output for debug
}

func Test2(t *testing.T) {
	candidates := []int{42, 26, 36, 38, 35, 41, 20, 47, 45, 23, 33, 39, 25, 43, 29, 31, 28, 48, 21, 46, 22, 30, 37, 32, 44, 40}
	res := combinationSum(candidates, 55)
	fmt.Printf("%+v\n", res) // output for debug
}
