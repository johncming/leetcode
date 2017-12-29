package leetcode

func maxAreaOfIsland(grid [][]int) int {
	current := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res := bfs(&grid, i, j) // init == 1
				if res > current {
					current = res
				}
			}
		}
	}

	return current
}

type item struct {
	m    *[][]int
	i, j int
}

func (i item) Left() (int, int, bool) {
	if i.j-1 < 0 {
		return 0, 0, false
	}
	if (*i.m)[i.i][i.j-1] == 1 {
		return i.i, i.j - 1, true
	} else {
		return i.i, i.j - 1, false
	}
}

func (i item) Right() (int, int, bool) {
	max := len((*i.m)[0]) - 1
	if i.j+1 > max {
		return 0, 0, false
	}
	if (*i.m)[i.i][i.j+1] == 1 {
		return i.i, i.j + 1, true
	} else {
		return i.i, i.j + 1, false
	}
}

func (i item) Up() (int, int, bool) {
	if i.i-1 < 0 {
		return 0, 0, false
	}
	if (*i.m)[i.i-1][i.j] == 1 {
		return i.i - 1, i.j, true
	} else {
		return i.i - 1, i.j, false
	}
}

func (i item) Down() (int, int, bool) {
	max := len(*i.m) - 1
	if i.i+1 > max {
		return 0, 0, false
	}
	if (*i.m)[i.i+1][i.j] == 1 {
		return i.i + 1, i.j, true
	} else {
		return i.i + 1, i.j, false
	}
}

type queue []item

func (q queue) Enqueue(v item) queue {
	return append(q, v)
}

func (q queue) Dequeue() (queue, item) {
	return q[1:], q[0]
}

func bfs(grid *[][]int, i, j int) int {
	var pos item

	res := 1
	q := make(queue, 0)
	q = q.Enqueue(item{grid, i, j})
	(*grid)[i][j] = 0
	for len(q) > 0 {
		q, pos = q.Dequeue()
		if i, j, ok := pos.Left(); ok {
			q = q.Enqueue(item{grid, i, j})
			(*grid)[i][j] = 0
			res++
		}
		if i, j, ok := pos.Right(); ok {
			q = q.Enqueue(item{grid, i, j})
			(*grid)[i][j] = 0
			res++
		}
		if i, j, ok := pos.Up(); ok {
			q = q.Enqueue(item{grid, i, j})
			(*grid)[i][j] = 0
			res++
		}
		if i, j, ok := pos.Down(); ok {
			q = q.Enqueue(item{grid, i, j})
			(*grid)[i][j] = 0
			res++
		}
	}

	return res
}
