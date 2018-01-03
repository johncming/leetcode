package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	long, short := nums1, nums2
	if len(nums2) > len(long) {
		long, short = nums2, nums1
	}

	positions := make(map[int]bool)
	for i, _ := range long {
		positions[i] = false
	}

	var res []int
L:
	for _, e1 := range short {
		for j, e2 := range long {
			if e1 == e2 && !positions[j] {
				res = append(res, e1)
				positions[j] = true
				continue L
			}
		}
	}

	return res
}
