package leetcode

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	nrows := len(grid)
	ncols := len(grid[0])

	var dp [][]int
	for i := 0; i < nrows; i++ {
		arr := make([]int, ncols)
		dp = append(dp, arr)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < ncols; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < nrows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < nrows; i++ {
		for j := 1; j < ncols; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[nrows-1][ncols-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
