package leetcode

func subsets(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	helper(&res, []int{}, nums, 0)
	return res
}

func helper(res *[][]int, list []int, nums []int, index int) {
	l := make([]int, len(list))
	copy(l, list)
	*res = append(*res, l)

	for i := index; i < len(nums); i++ {
		list = append(list, nums[i])
		helper(res, list, nums, i+1)
		list = list[:len(list)-1]
	}
}
