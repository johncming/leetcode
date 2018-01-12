package leetcode

func threeSumClosest(nums []int, target int) int {
	L := len(nums)

	diff := 1<<31 - 1
	result := 0
	for i := 0; i < L; i++ {
		for j := i + 1; j < L; j++ {
			for k := j + 1; k < L; k++ {
				sum := nums[i] + nums[j] + nums[k]
				tmp := sum - target

				switch {
				case tmp == 0:
					return sum
				case tmp > 0:
					if tmp < diff {
						diff = tmp
						result = sum
					}
				case tmp < 0:
					if -tmp < diff {
						diff = -tmp
						result = sum
					}
				}
			}
		}
	}

	return result
}
