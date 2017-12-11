package leetcode

func lengthOfLongestSubstring(s string) int {
	var longest []rune

	var tmp []rune
	for _, ch := range s {
		for i, x := range tmp {
			if ch == x {
				if len(longest) < len(tmp) {
					longest = tmp
				}
				tmp2 := make([]rune, len(tmp))
				copy(tmp2, tmp)
				tmp = tmp2[i+1:]
				break
			}
		}
		tmp = append(tmp, ch)
	}

	if len(tmp) > len(longest) {
		longest = tmp
	}

	return len(longest)
}
