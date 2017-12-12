package leetcode

func longestPalindrome(s string) string {
	var longest []rune

	for start, _ := range s {
		tmp := make([]rune, 0)
		for _, ch := range s[start:] {
			if isPalindrome(tmp) {
				if len(longest) < len(tmp) {
					longest = tmp
				}
				tmp2 := make([]rune, len(tmp))
				copy(tmp2, tmp)
				tmp = tmp2
			}
			tmp = append(tmp, ch)
		}
		if isPalindrome(tmp) {
			if len(longest) < len(tmp) {
				longest = tmp
			}
			tmp2 := make([]rune, len(tmp))
			copy(tmp2, tmp)
			tmp = tmp2
		}
	}

	return string(longest)
}

func isPalindrome(s []rune) bool {

	if len(s) == 0 {
		return false
	}

	start := 0
	end := len(s) - 1

	for start <= end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			return false
		}
	}

	return true
}
