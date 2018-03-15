package leetcode

func generateParenthesis(n int) []string {
	var res []string
	if n == 0 {
		return res
	}
	helper(&res, "", n, n)
	return res
}

func helper(res *[]string, temp string, left, right int) {
	if left > right {
		return
	}

	if left == 0 && right == 0 {
		*res = append(*res, temp)
		return
	}

	if left > 0 {
		helper(res, temp+"(", left-1, right)
	}

	if right > 0 {
		helper(res, temp+")", left, right-1)
	}

}
