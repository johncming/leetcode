package leetcode

import (
	"fmt"
	"log"
)

func summaryRanges(nums []int) []string {

	var res []string
	var stop int

	maxIndex := len(nums) - 1

	start := 0

	for i, n := range nums {

		if i+1 > maxIndex {
			log.Println("stage 1")
			stop = maxIndex
			if start == stop {
				res = append(res, fmt.Sprintf("%d", nums[start]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[stop]))
			}
			break
		}

		if n+1 != nums[i+1] {
			log.Println("stage 2")
			stop = i
			if start == stop {
				res = append(res, fmt.Sprintf("%d", nums[start]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[stop]))
			}
			start = i + 1
		}

	}

	return res

}
