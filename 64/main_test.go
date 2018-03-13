package leetcode

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	grid := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}

	res := minPathSum(grid)
	fmt.Printf("%+v\n", res) // output for debug
}

func Test2(t *testing.T) {
	grid := [][]int{
		[]int{1, 2, 5},
		[]int{3, 2, 1},
	}

	res := minPathSum(grid)
	fmt.Printf("%+v\n", res) // output for debug
}
