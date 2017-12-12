package leetcode

func romanToInt(s string) int {
	var res int
	i := 0
	for i < len(s) {
		switch s[i] {
		case 'M':
			res += 1000
		case 'D':
			res += 500
		case 'C':
			if i+1 >= len(s) {
				res += 100
				break
			}
			switch s[i+1] {
			case 'D':
				res += 400
				i++
			case 'M':
				res += 900
				i++
			default:
				res += 100
			}
		case 'L':
			res += 50
		case 'X':
			if i+1 >= len(s) {
				res += 10
				break
			}
			switch s[i+1] {
			case 'C':
				res += 90
				i++
			case 'L':
				res += 40
				i++
			default:
				res += 10
			}
		case 'V':
			res += 5
		case 'I':
			if i+1 >= len(s) {
				res += 1
				break
			}
			switch s[i+1] {
			case 'X':
				res += 9
				i++
			case 'V':
				res += 4
				i++
			default:
				res += 1
			}

		}
		i++
	}

	return res
}
