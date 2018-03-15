package leetcode

type stack []rune

func (s stack) push(r rune) stack {
	s = append(s, r)
	return s
}

func (s stack) pop() (stack, rune) {
	return s[:len(s)-1], s[len(s)-1]
}

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}

	var S stack

	for _, c := range s {
		switch c {
		case '(':
			S = S.push(')')
		case '[':
			S = S.push(']')
		case '{':
			S = S.push('}')
		default:
			if len(S) == 0 {
				return false
			}
			var c1 rune
			S, c1 = S.pop()
			if c1 != c {
				return false
			}
		}
	}

	return len(S) == 0
}
