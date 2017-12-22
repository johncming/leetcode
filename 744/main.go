package leetcode

func nextGreatestLetter(letters []byte, target byte) byte {
	lastIndex := len(letters) - 1
	first := letters[0]
	letters = append(letters, first)

	var res byte
	for i, le := range letters {
		if le <= target {
			if i > lastIndex {
				res = le
				break
			}
			continue
		} else {
			res = le
			break
		}
	}

	return res
}
