package leetcode

var result [][]int

func combinationSum(candidates []int, target int) [][]int {
	result = [][]int{}
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
		c := candidates[i]
		temp = append(temp, c)
		dfs(candidates, temp, reminder-c, i)
		temp = temp[:len(temp)-1]
	}
}
