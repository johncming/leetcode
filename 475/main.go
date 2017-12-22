package leetcode

import (
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	var ret int

	sort.Ints(heaters)

	for _, house := range houses {
		index := sort.SearchInts(heaters, house)
		if index >= len(heaters) {
			diff := house - heaters[len(heaters)-1]
			if diff > ret {
				ret = diff
			}
		} else if index == 0 {
			diff := heaters[0] - house
			if diff > ret {
				ret = diff
			}
		} else {
			diff1 := house - heaters[index-1]
			diff2 := heaters[index] - house

			if diff1 < diff2 {
				diff2 = diff1
			}
			if diff2 > ret {
				ret = diff2
			}
		}
	}

	return ret
}
