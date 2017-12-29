package leetcode

func islandPerimeter(grid [][]int) int {
	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				cnt += 4
				if j > 0 && grid[i][j-1] == 1 {
					cnt -= 2
				}
				if i > 0 && grid[i-1][j] == 1 {
					cnt -= 2
				}
			}
		}
	}
	return cnt
}
