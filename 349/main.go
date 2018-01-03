package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	var tmp []int
	for _, e1 := range nums1 {
		for _, e2 := range nums2 {
			if e1 == e2 {
				tmp = append(tmp, e1)
			}
		}
	}

	cache := make(map[int]struct{})
	for _, item := range tmp {
		cache[item] = struct{}{}
	}
	var res []int
	for item, _ := range cache {
		res = append(res, item)
	}
	return res
}
