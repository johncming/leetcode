package leetcode

func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 || digits == "1" {
		return res
	}
	helper(&res, "", digits)
	return res
}

func helper(res *[]string, temp string, digits string) {
	// edge
	if len(digits) == 0 {
		*res = append(*res, temp)
		return
	}

	switch digits[0] {
	case '2':
		for _, c := range []string{"a", "b", "c"} {
			helper(res, temp+c, digits[1:])
		}
	case '3':
		for _, c := range []string{"d", "e", "f"} {
			helper(res, temp+c, digits[1:])
		}
	case '4':
		for _, c := range []string{"g", "h", "i"} {
			helper(res, temp+c, digits[1:])
		}
	case '5':
		for _, c := range []string{"j", "k", "l"} {
			helper(res, temp+c, digits[1:])
		}
	case '6':
		for _, c := range []string{"m", "n", "o"} {
			helper(res, temp+c, digits[1:])
		}
	case '7':
		for _, c := range []string{"p", "q", "r", "s"} {
			helper(res, temp+c, digits[1:])
		}
	case '8':
		for _, c := range []string{"t", "u", "v"} {
			helper(res, temp+c, digits[1:])
		}
	case '9':
		for _, c := range []string{"w", "x", "y", "z"} {
			helper(res, temp+c, digits[1:])
		}
	default: // 1 and 0
	}
}
