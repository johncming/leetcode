package leetcode

import (
	"sort"
)

var result [][]int

func combinationSum2(candidates []int, target int) [][]int {
	result = [][]int{}
	sort.Ints(candidates)
	dfs(candidates, []int{}, target, 0)
	return result
}

func dfs(candidates []int, temp []int, reminder int, index int) {
	if reminder == 0 {
		t := make([]int, len(temp))
		copy(t, temp)
		result = append(result, t)
		return
	}

	if reminder < 0 {
		return
	}

	for i := index; i < len(candidates); i++ {
		if i != index && candidates[i] == candidates[i-1] {
			continue
		}
		c := candidates[i]
		temp = append(temp, c)
		dfs(candidates, temp, reminder-c, i+1)
		temp = temp[:len(temp)-1]
	}
}
