package leetcode

func islandPerimeter(grid [][]int) int {
	cnt := 0
	neighbour := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				cnt++
				if j+1 < len(grid[0]) && grid[i][j+1] == 1 {
					neighbour++
				}
				if i+1 < len(grid) && grid[i+1][j] == 1 {
					neighbour++
				}
			}
		}
	}
	// 每增加一个邻居，少两条边
	return cnt*4 - neighbour*2
}
