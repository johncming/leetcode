package leetcode

func removeElement(nums []int, val int) int {

	var index, another []int

	for i, n := range nums {
		if n == val {
			index = append(index, i)
		} else {
			another = append(another, i)
		}
	}

	lastAnother := len(another) - 1 // index

	if lastAnother >= 0 {

		for _, i := range index {
			if lastAnother < 0 {
				break
			}
			nums[i] = nums[another[lastAnother]]
			lastAnother--
		}
	}

	return len(another)

}
