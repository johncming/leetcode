package leetcode

func fourSumCount(A []int, B []int, C []int, D []int) int {

	cache := make(map[int]int)
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			sum := C[i] + D[j]
			if v, ok := cache[sum]; !ok {
				cache[sum] = 1
			} else {
				cache[sum] = v + 1
			}
		}
	}

	res := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			sum := A[i] + B[j]
			if v, ok := cache[-sum]; ok {
				res += v
			}
		}
	}

	return res
}
