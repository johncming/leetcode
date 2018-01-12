package leetcode

func longestPalindrome(s string) string {
	hash := make(map[rune][]int)

	for i, c := range s {
		hash[c] = append(hash[c], i)
	}

	longest := ""
	for _, v := range hash {
		tmp := ""
		for i := 0; i < len(v); i++ {
			for j := len(v) - 1; j >= i; j-- {
				if isPalindrome(s, v[i], v[j]) {
					tmp = s[v[i] : v[j]+1]
					if len(tmp) > len(longest) {
						longest = tmp
					}
				}
			}
		}
	}

	return longest
}

func isPalindrome(s string, left, right int) bool {
	for i, j := left, right; i <= j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
