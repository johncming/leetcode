package leetcode

var result [][]int

func combine(n int, k int) [][]int {
	result = [][]int{}

	dfs(k, 1, n, []int{})
	return result
}

func dfs(remainder int, candidate, n int, temp []int) {
	// condition
	if remainder == 0 {
		t := make([]int, len(temp))
		copy(t, temp)
		result = append(result, t)
		return
	}

	for i := candidate; i <= n; i++ {
		temp = append(temp, i)
		dfs(remainder-1, i+1, n, temp)
		temp = temp[:len(temp)-1]
	}
}
