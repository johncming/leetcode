package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	ss := make([][]byte, numRows)

	for i := 0; i < len(s); i++ {
		// 斜处没有head和tail，所以减2
		index := i % (2*numRows - 2)
		if index >= numRows {
			// 比如numRows=3, index=3, 则2*3-2=4, 4-3=1，所以放在第二行
			// 比如numRows=4, index=4, 则2*4-2=6, 6-4=2，以此类推
			index = 2*numRows - 2 - index
		}
		ss[index] = append(ss[index], s[i])
	}

	for i := 1; i < numRows; i++ {
		ss[0] = append(ss[0], ss[i]...)
	}
	return string(ss[0])
}
