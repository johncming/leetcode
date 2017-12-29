package leetcode

func numIslands(grid [][]byte) int {
	length := len(grid)
	res := 0
	for i := 0; i < length; i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0x31 {
				res++
				bfs(&grid, i, j)
			}
		}
	}
	return res
}

type Item struct {
	m    *[][]byte
	i, j int
}

func (i Item) Left() (int, int, bool) {
	if i.j-1 < 0 {
		return 0, 0, false
	}
	if (*i.m)[i.i][i.j-1] == 0x31 {
		return i.i, i.j - 1, true
	} else {
		return 0, 0, false
	}
}

func (i Item) Right() (int, int, bool) {
	max := len((*i.m)[0]) - 1
	if i.j+1 > max {
		return 0, 0, false
	}

	if (*i.m)[i.i][i.j+1] == 0x31 {
		return i.i, i.j + 1, true
	} else {
		return 0, 0, false
	}
}

func (i Item) Up() (int, int, bool) {
	if i.i-1 < 0 {
		return 0, 0, false
	}
	if (*i.m)[i.i-1][i.j] == 0x31 {
		return i.i - 1, i.j, true
	} else {
		return 0, 0, false
	}
}

func (i Item) Down() (int, int, bool) {
	max := len(*i.m) - 1
	if i.i+1 > max {
		return 0, 0, false
	}
	if (*i.m)[i.i+1][i.j] == 0x31 {
		return i.i + 1, i.j, true
	} else {
		return 0, 0, false
	}
}

type Queue []Item

func (q Queue) Append(v Item) Queue {
	return append(q, v)
}

func (q Queue) PopLeft() (Queue, Item) {
	return q[1:], q[0]
}

func bfs(grid *[][]byte, i, j int) {
	q := Item{grid, i, j}
	queue := make(Queue, 0)
	queue = queue.Append(q)
	for len(queue) > 0 {
		queue, q = queue.PopLeft()
		if i, j, ok := q.Left(); ok {
			queue = queue.Append(Item{grid, i, j})
			(*grid)[i][j] = 0
		}
		if i, j, ok := q.Right(); ok {
			queue = queue.Append(Item{grid, i, j})
			(*grid)[i][j] = 0
		}
		if i, j, ok := q.Up(); ok {
			queue = queue.Append(Item{grid, i, j})
			(*grid)[i][j] = 0
		}
		if i, j, ok := q.Down(); ok {
			queue = queue.Append(Item{grid, i, j})
			(*grid)[i][j] = 0
		}
	}
}
