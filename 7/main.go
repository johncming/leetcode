package leetcode

import (
	"fmt"
)

func reverse(x int) int {
	max := 1<<31 - 1

	sx := fmt.Sprintf("%d", x)

	var ss []byte
	for i := len(sx) - 1; i >= 0; i-- {
		ss = append(ss, sx[i])
	}

	// negative
	var st []byte
	var negative bool
	if ss[len(ss)-1] == '-' {
		negative = true
		ss = ss[:len(ss)-1]
		st = append(st, '-')
		st = append(st, ss...)
	}

	var sss string
	if negative {
		sss = string(st)
	} else {
		sss = string(ss)
	}

	var y int
	fmt.Sscanf(sss, "%d", &y)
	if y > max || y < -max {
		return 0
	}
	return y
}
