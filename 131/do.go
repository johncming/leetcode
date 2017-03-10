package leetcode

import (
	"strings"
)

func partition(s string) [][]string {
	var res [][]string

	res = append(res, strings.Split(s, ""))

	last := len(s) - 1
	for i, _ := range s {
		var _res []string
		next := i + 1

		for {
			if next > last {
				break
			}

			sub := s[i:next]

			if isPalindrome(sub) {
				_res = append(_res, sub)
			}

			next++
		}
		res = append(res, _res)
	}

	return res
}

func isPalindrome(s string) (res bool) {

	right := len(s) - 1
	left := 0

	for {
		if left >= right {
			res = true
			break
		}
		if s[left] == s[right] {
			left++
			right--
			continue
		} else {
			return
		}
	}

	return
}
